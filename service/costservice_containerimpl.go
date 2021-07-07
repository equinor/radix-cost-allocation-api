package service

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/equinor/radix-common/utils"
	"github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/repository"
)

func ContainerMissingNodeError(containerId string) error {
	return fmt.Errorf("container %s, node is nil", containerId)
}

func ContainerMissingNodePoolIdError(containerId string) error {
	return fmt.Errorf("container %s, node pool Id is nil", containerId)
}

func CurrencyMismatchError(expected, actual string) error {
	return fmt.Errorf("expected currency %s but got %s", expected, actual)
}

type containerTotalCost struct {
	container *models.ContainerDto
	cost
}

type cost struct {
	value    float64
	currency string
}

type containerCost struct {
	containerID string
	cost
}

type containerResourceUsage struct {
	containerID         string
	cpuMillicoreSeconds float64
	memoryBytesSeconds  float64
}

type nodePoolCostAllocatedResources struct {
	cost                float64
	currency            string
	cpuMillicoreSeconds float64
	memoryBytesSeconds  float64
	containerResources  []containerResourceUsage
}

type containerCostService struct {
	repo                   repository.Repository
	applicationExcludeList []string
}

func NewContainerCostService(repo repository.Repository, applicationExcludeList []string) CostService {
	return &containerCostService{
		repo:                   repo,
		applicationExcludeList: applicationExcludeList,
	}
}

func (s *containerCostService) GetCostForPeriod(from, to time.Time) (*models.ApplicationCostSet, error) {
	applicationCostList, err := s.getApplicationCostList(from, to)
	if err != nil {
		return nil, err
	}

	applicationCost := models.ApplicationCostSet{
		From:             from,
		To:               to,
		ApplicationCosts: applicationCostList,
	}
	return &applicationCost, nil
}

func (s *containerCostService) GetFutureCost(appName string) (*models.ApplicationCost, error) {
	to := time.Now()
	from := to.Add(-24 * time.Hour)
	applicationCostList, err := s.getApplicationCostList(from, to)
	if err != nil {
		return nil, err
	}

	for _, appCost := range applicationCostList {
		if strings.EqualFold(appCost.Name, appName) {
			appCost := appCost
			appCost.Cost *= 30 // Multiply cost for last 24 hours with 30 days
			return &appCost, nil
		}
	}

	return &models.ApplicationCost{Name: appName}, nil
}

func (s *containerCostService) getApplicationCostList(from, to time.Time) ([]models.ApplicationCost, error) {
	nodePools, err := s.repo.GetNodePools()
	if err != nil {
		return nil, err
	}

	nodePoolCost, err := s.repo.GetNodePoolCost()
	if err != nil {
		return nil, err
	}

	containers, err := s.repo.GetContainers(from, to)
	if err != nil {
		return nil, err
	}

	containers = excludeApplicationNames(containers, s.applicationExcludeList)
	poolCost := buildNodePoolCost(from, to, nodePools, nodePoolCost)
	containerCostList, err := calculateContainerCost(poolCost, containers)
	if err != nil {
		return nil, err
	}

	aggregatedContainerCost, err := aggregateContainerCost(containerCostList, containers)
	if err != nil {
		return nil, err
	}

	applicationCostList := buildApplicationCostList(aggregatedContainerCost)
	return applicationCostList, nil
}

func buildApplicationCostList(containerTotalCostList []containerTotalCost) []models.ApplicationCost {
	appCostIndex := make(map[string]int)
	var applicationCostList []models.ApplicationCost
	wbsCodes := make(map[string]time.Time)

	for _, c := range containerTotalCostList {
		if idx, ok := appCostIndex[c.container.ApplicationName]; !ok {
			appCost := models.ApplicationCost{Name: c.container.ApplicationName, Cost: c.value, Currency: c.currency, WBS: c.container.WBS}
			applicationCostList = append(applicationCostList, appCost)
			appCostIndex[c.container.ApplicationName] = len(applicationCostList) - 1
			wbsCodes[c.container.ApplicationName] = c.container.LastKnownRunningAt
		} else {
			applicationCostList[idx].Cost += c.value

			if c.container.LastKnownRunningAt.After(wbsCodes[c.container.ApplicationName]) {
				wbsCodes[c.container.ApplicationName] = c.container.LastKnownRunningAt
				applicationCostList[idx].WBS = c.container.WBS
			}
		}
	}

	return applicationCostList
}

func excludeApplicationNames(containers []models.ContainerDto, applicationNames []string) []models.ContainerDto {
	if len(applicationNames) == 0 {
		return containers
	}

	var i int
	for _, c := range containers {
		if utils.ContainsString(applicationNames, c.ApplicationName) {
			continue
		}
		containers[i] = c
		i++
	}
	return containers[:i]
}

func aggregateContainerCost(containerCostList []containerCost, containers []models.ContainerDto) ([]containerTotalCost, error) {
	indexMap := make(map[string]int)
	containerTotalCostList := make([]containerTotalCost, len(containers))

	for i, c := range containers {
		c := c
		indexMap[c.ContainerId] = i
		containerTotalCostList[i] = containerTotalCost{container: &c}
	}

	for _, cost := range containerCostList {
		if i, ok := indexMap[cost.containerID]; ok {
			if containerTotalCostList[i].currency == "" {
				containerTotalCostList[i].currency = strings.ToUpper(cost.currency)
			}

			if containerTotalCostList[i].currency != strings.ToUpper(cost.currency) {
				return nil, CurrencyMismatchError(containerTotalCostList[i].currency, cost.currency)
			}

			containerTotalCostList[i].value += cost.value
		}
	}

	return containerTotalCostList, nil
}

func calculateContainerCost(poolCosts []models.NodePoolCostDto, containers []models.ContainerDto) ([]containerCost, error) {
	var containerCostList []containerCost

	for _, cost := range poolCosts {
		nodePoolCostResource, err := getAllocatedResourcesForNodePoolCost(cost, containers)
		if err != nil {
			return nil, err
		}

		nodePoolContainerCost := calculateNodePoolContainerResourceCost(nodePoolCostResource)
		containerCostList = append(containerCostList, nodePoolContainerCost...)
	}

	return containerCostList, nil
}

func calculateNodePoolContainerResourceCost(nodePoolCostResource nodePoolCostAllocatedResources) (containerCostList []containerCost) {
	for _, resource := range nodePoolCostResource.containerResources {
		containerCostList = append(containerCostList, containerCost{
			containerID: resource.containerID,
			cost: cost{
				value: calculateContainerResourceCost(
					nodePoolCostResource.cpuMillicoreSeconds,
					nodePoolCostResource.memoryBytesSeconds,
					resource.cpuMillicoreSeconds,
					resource.memoryBytesSeconds,
					nodePoolCostResource.cost),
				currency: nodePoolCostResource.currency,
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

func getAllocatedResourcesForNodePoolCost(cost models.NodePoolCostDto, containers []models.ContainerDto) (nodePoolCostResource nodePoolCostAllocatedResources, err error) {
	var cpuSec, memSec float64
	nodePoolCostResource.cost = cost.Cost
	nodePoolCostResource.currency = cost.Currency

	for _, cont := range containers {
		contCpuSec, contMemSec, callErr := getContainerResourcesUsageInNodePoolCost(cost, cont)
		if callErr != nil {
			err = callErr
			return
		}

		if contCpuSec > 0 || contMemSec > 0 {
			cru := containerResourceUsage{containerID: cont.ContainerId, cpuMillicoreSeconds: contCpuSec, memoryBytesSeconds: contMemSec}
			nodePoolCostResource.containerResources = append(nodePoolCostResource.containerResources, cru)
			cpuSec += contCpuSec
			memSec += contMemSec
		}
	}

	nodePoolCostResource.cpuMillicoreSeconds = cpuSec
	nodePoolCostResource.memoryBytesSeconds = memSec

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
	sort.Sort(sortByFromAndTo(cost))

	for i := 0; i < len(cost)-1; i++ {
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
	adjustedCost = append(adjustedCost, cost[len(cost)-1])

	// Adjust FromDate of first entry to range from date if FromDate>from
	if adjustedCost[0].FromDate.After(from) {
		adjustedCost[0] = adjustCostPeriod(adjustedCost[0], from, adjustedCost[0].ToDate)
	}

	// Adjust ToDate of last entry to range to date if ToDate<to
	lastIdx := len(adjustedCost) - 1
	if adjustedCost[lastIdx].ToDate.Before(to) {
		adjustedCost[lastIdx] = adjustCostPeriod(adjustedCost[lastIdx], adjustedCost[lastIdx].FromDate, to)
	}

	// Remove cost entries outside from and to range
	var idx int
	for _, c := range adjustedCost {
		if !isCostInsideRange(from, to, c) {
			continue
		}
		adjustedCost[idx] = c
		idx++
	}
	adjustedCost = adjustedCost[:idx]

	// Adjust FromDate and ToDate in first and last cost item to match from and to range
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

// NodePoolCostDto sorter - sorts by FromDate, and for entries where FromDate is equal we sort by ToDate
// Example entries
//   |---------|
//        |-------|
//        |----------------|
//             |----|
type sortByFromAndTo []models.NodePoolCostDto

func (c sortByFromAndTo) Len() int { return len(c) }
func (c sortByFromAndTo) Less(i, j int) bool {

	if c[i].FromDate.Before(c[j].FromDate) {
		return true
	} else if c[i].FromDate.After(c[j].FromDate) {
		return false
	}

	return c[i].ToDate.Before(c[j].ToDate)
}
func (c sortByFromAndTo) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
