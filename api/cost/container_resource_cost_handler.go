package cost

import (
	"sort"
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/repository"
)

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

	buildNodePoolCost(from, to, nodePools, nodePoolCost)

	return nil, nil
}

func buildNodePoolCost(from, to time.Time, nodePools []models.NodePoolDto, nodePoolCost []models.NodePoolCostDto) {
	for _, pool := range nodePools {
		cost := filterNodePoolCostByPoolId(nodePoolCost, pool.Id)
		cost = adjustNodePoolCostTimeRange(from, to, cost)
	}

	return
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
	cost.Cost = int32(newDuration / currentDuration * float64(cost.Cost))

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
