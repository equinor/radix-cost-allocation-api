package cost

import (
	"fmt"
	"sort"
	"strings"
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/repository"
)

type ContainerTotalCost struct {
	Container *models.ContainerDto
	Cost
}

type Cost struct {
	Value    float64
	Currency string
}

type ContainerCost struct {
	ContainerId string
	Cost
}

type ContainerResourceUsage struct {
	ContainerId         string
	CPUMillicoreSeconds float64
	MemoryBytesSeconds  float64
}

type NodePoolCostAllocatedResources struct {
	Cost                float64
	Currency            string
	CPUMillicoreSeconds float64
	MemoryBytesSeconds  float64
	ContainerResources  []ContainerResourceUsage
}

func ContainerMissingNodeError(containerId string) error {
	return fmt.Errorf("container %s, node is nil", containerId)
}

func ContainerMissingNodePoolIdError(containerId string) error {
	return fmt.Errorf("container %s, node pool Id is nil", containerId)
}

func CurrencyMismatchError(expected, actual string) error {
	return fmt.Errorf("expected currency %s but got %s", expected, actual)
}

type containerResourceCostHandler struct {
	repo repository.Repository
}

func NewContainerResourceCostHandler(repo repository.Repository) *containerResourceCostHandler {
	return &containerResourceCostHandler{
		repo: repo,
	}
}

func (h *containerResourceCostHandler) GetTotalCost(from time.Time, to time.Time) (*costModels.ApplicationCostSet, error) {
	nodePools, err := h.repo.GetNodePools()
	if err != nil {
		return nil, err
	}

	nodePoolCost, err := h.repo.GetNodePoolCost(from, to)
	if err != nil {
		return nil, err
	}

	containers, err := h.repo.GetContainers(from, to)

	poolCost := buildNodePoolCost(from, to, nodePools, nodePoolCost)
	containerCost, err := calculateContainerCost(poolCost, containers)
	if err != nil {
		return nil, err
	}

	aggregatedContainerCost, err := aggregateContainerCost(containerCost, containers)
	if err != nil {
		return nil, err
	}

	fmt.Println(aggregatedContainerCost)

	return nil, nil
}

func aggregateContainerCost(containerCost []ContainerCost, containers []models.ContainerDto) ([]ContainerTotalCost, error) {
	indexMap := make(map[string]int)
	containerCostList := make([]ContainerTotalCost, len(containers))

	for i, c := range containers {
		c := c
		indexMap[c.ContainerId] = i
		containerCostList[i] = ContainerTotalCost{Container: &c}
	}

	for _, cost := range containerCost {
		if i, ok := indexMap[cost.ContainerId]; ok {
			if containerCostList[i].Currency == "" {
				containerCostList[i].Currency = strings.ToUpper(cost.Currency)
			}

			if containerCostList[i].Currency != strings.ToUpper(cost.Currency) {
				return nil, CurrencyMismatchError(containerCostList[i].Currency, cost.Currency)
			}

			containerCostList[i].Value += cost.Value
		}
	}

	return containerCostList, nil
}

func calculateContainerCost(poolCosts []models.NodePoolCostDto, containers []models.ContainerDto) ([]ContainerCost, error) {
	var containerCost []ContainerCost

	for _, cost := range poolCosts {
		nodePoolCostResource, err := getAllocatedResourcesForNodePoolCost(cost, containers)
		if err != nil {
			return nil, err
		}

		nodePoolContainerCost := calculateNodePoolContainerResourceCost(nodePoolCostResource)
		containerCost = append(containerCost, nodePoolContainerCost...)
	}

	return containerCost, nil
}

func calculateNodePoolContainerResourceCost(nodePoolCostResource NodePoolCostAllocatedResources) (containerCost []ContainerCost) {
	for _, resource := range nodePoolCostResource.ContainerResources {
		containerCost = append(containerCost, ContainerCost{
			ContainerId: resource.ContainerId,
			Cost: Cost{
				Value: calculateContainerResourceCost(
					nodePoolCostResource.CPUMillicoreSeconds,
					nodePoolCostResource.MemoryBytesSeconds,
					resource.CPUMillicoreSeconds,
					resource.MemoryBytesSeconds,
					nodePoolCostResource.Cost),
				Currency: nodePoolCostResource.Currency,
			},
		})
	}

	return
}

func calculateContainerResourceCost(nodepoolCpuSeconds, nodepoolMemorySeconds, containerCpuSeconds, containerMemorySeconds, nodepoolCost float64) float64 {
	cpuCost := containerCpuSeconds / nodepoolCpuSeconds * nodepoolCost / 2
	memCost := containerMemorySeconds / nodepoolMemorySeconds * nodepoolCost / 2
	return cpuCost + memCost
}

func getAllocatedResourcesForNodePoolCost(cost models.NodePoolCostDto, containers []models.ContainerDto) (nodePoolCostResource NodePoolCostAllocatedResources, err error) {
	var cpuSec, memSec float64
	nodePoolCostResource.Cost = cost.Cost
	nodePoolCostResource.Currency = cost.Currency

	for _, cont := range containers {
		contCpuSec, contMemSec, callErr := getContainerResourcesUsageInNodePoolCost(cost, cont)
		if callErr != nil {
			err = callErr
			return
		}

		if contCpuSec > 0 || contMemSec > 0 {
			containerResourceUsage := ContainerResourceUsage{ContainerId: cont.ContainerId, CPUMillicoreSeconds: contCpuSec, MemoryBytesSeconds: contMemSec}
			nodePoolCostResource.ContainerResources = append(nodePoolCostResource.ContainerResources, containerResourceUsage)
			cpuSec += contCpuSec
			memSec += contMemSec
		}
	}

	nodePoolCostResource.CPUMillicoreSeconds = cpuSec
	nodePoolCostResource.MemoryBytesSeconds = memSec

	return
}

func getContainerResourcesUsageInNodePoolCost(cost models.NodePoolCostDto, container models.ContainerDto) (cpuSec float64, memSec float64, err error) {
	if container.Node == nil {
		err = ContainerMissingNodeError(container.ContainerId)
		return
	}

	if container.Node.NodePoolId == nil {
		err = ContainerMissingNodePoolIdError(container.ContainerId)
		return
	}

	if *container.Node.NodePoolId == cost.NodePoolId {
		duration := getContainerDurationInNodePoolCost(cost, container)
		cpuSec = duration.Seconds() * float64(container.CpuRequestedMillicores)
		memSec = duration.Seconds() * float64(container.MemoryRequestedBytes)
	}

	return
}

func getContainerDurationInNodePoolCost(cost models.NodePoolCostDto, container models.ContainerDto) time.Duration {
	if isContainerRunningInNodePoolCost(cost, container) {
		duration := container.LastKnownRunningAt.Sub(container.StartedAt)

		if container.StartedAt.Before(cost.FromDate) {
			duration -= cost.FromDate.Sub(container.StartedAt)
		}

		if container.LastKnownRunningAt.After(cost.ToDate) {
			duration -= container.LastKnownRunningAt.Sub(cost.ToDate)
		}

		return duration
	}

	return 0
}

func isContainerRunningInNodePoolCost(cost models.NodePoolCostDto, container models.ContainerDto) bool {
	return container.LastKnownRunningAt.After(cost.FromDate) && container.StartedAt.Before(cost.ToDate)
}

func buildNodePoolCost(from, to time.Time, nodePools []models.NodePoolDto, nodePoolCost []models.NodePoolCostDto) []models.NodePoolCostDto {
	var poolCosts []models.NodePoolCostDto

	for _, pool := range nodePools {
		cost := filterNodePoolCostByPoolId(nodePoolCost, pool.Id)
		cost = adjustNodePoolCostTimeRange(from, to, cost)
		poolCosts = append(poolCosts, cost...)
	}

	return poolCosts
}

func filterNodePoolCostByPoolId(cost []models.NodePoolCostDto, poolId int32) []models.NodePoolCostDto {
	var filteredCost []models.NodePoolCostDto

	for _, c := range cost {
		if c.NodePoolId == poolId {
			filteredCost = append(filteredCost, c)
		}
	}

	return filteredCost
}

func adjustNodePoolCostTimeRange(from, to time.Time, cost []models.NodePoolCostDto) []models.NodePoolCostDto {
	if len(cost) == 0 {
		return cost
	}

	adjustedCost := make([]models.NodePoolCostDto, 0, len(cost))
	sort.Sort(SortByFromAndTo(cost))

	for i := 0; i < len(cost)-1; i++ {
		if !isCostInsideRange(from, to, cost[i]) {
			continue
		}

		if isCostEncapsulated(cost[i], cost[i+1]) {
			continue
		}

		if !isCostConnected(cost[i], cost[i+1]) {
			adjustedCost = append(adjustedCost, adjustCostPeriod(cost[i], cost[i].FromDate, cost[i+1].FromDate))
			continue
		}

		adjustedCost = append(adjustedCost, cost[i])
	}

	// Add the last cost from source
	if isCostInsideRange(from, to, cost[len(cost)-1]) {
		adjustedCost = append(adjustedCost, cost[len(cost)-1])
	}

	// Adjust FromDate and ToDate in first and last cost item to match to and from range
	if len(adjustedCost) > 0 {
		adjustedCost[0] = adjustCostPeriod(adjustedCost[0], from, adjustedCost[0].ToDate)
		lastIdx := len(adjustedCost) - 1
		adjustedCost[lastIdx] = adjustCostPeriod(adjustedCost[lastIdx], adjustedCost[lastIdx].FromDate, to)
	}

	return adjustedCost
}

func adjustCostPeriod(cost models.NodePoolCostDto, newFromDate, newToDate time.Time) models.NodePoolCostDto {
	currentDuration := cost.ToDate.Sub(cost.FromDate).Seconds()
	newDuration := newToDate.Sub(newFromDate).Seconds()

	cost.FromDate = newFromDate
	cost.ToDate = newToDate
	cost.Cost = newDuration / currentDuration * cost.Cost

	return cost
}

func isCostInsideRange(rangeFrom, rangeTo time.Time, cost models.NodePoolCostDto) bool {
	return cost.ToDate.After(rangeFrom) && cost.FromDate.Before(rangeTo)
}

func isCostConnected(first, second models.NodePoolCostDto) bool {
	return first.ToDate.Equal(second.FromDate)
}

func isCostEncapsulated(first, second models.NodePoolCostDto) bool {
	fromGte := first.FromDate.Equal(second.FromDate) || first.FromDate.After(second.FromDate)
	toLte := first.ToDate.Equal(second.ToDate) || first.ToDate.Before(second.ToDate)
	return fromGte && toLte
}

type SortByFromAndTo []models.NodePoolCostDto

func (c SortByFromAndTo) Len() int { return len(c) }
func (c SortByFromAndTo) Less(i, j int) bool {

	if c[i].FromDate.Before(c[j].FromDate) {
		return true
	} else if c[i].FromDate.After(c[j].FromDate) {
		return false
	}

	return c[i].ToDate.Before(c[j].ToDate)
}
func (c SortByFromAndTo) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
