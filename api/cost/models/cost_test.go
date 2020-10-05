package cost_models_test

import (
	"math"
	"sort"
	"strconv"
	"testing"
	"time"

	cost_models "github.com/equinor/radix-cost-allocation-api/api/cost/models"

	"github.com/stretchr/testify/assert"
)

const subscriptionCost = 100000
const subscriptionCostCurrency = "NOK"

func Test_cost_all_app_equal_requested(t *testing.T) {
	runs := getTestRuns()
	cost := cost_models.NewApplicationCostSet(time.Now().Add(-1), time.Now(), runs, 0, "")

	assert.Equal(t, cost.ApplicationCosts[0].CostPercentageByCPU, cost.ApplicationCosts[1].CostPercentageByCPU)
	assert.Equal(t, cost.ApplicationCosts[1].CostPercentageByCPU, cost.ApplicationCosts[2].CostPercentageByCPU)
	assert.Equal(t, cost.ApplicationCosts[2].CostPercentageByCPU, cost.ApplicationCosts[3].CostPercentageByCPU)

	assert.Equal(t, cost.ApplicationCosts[0].CostPercentageByMemory, cost.ApplicationCosts[1].CostPercentageByMemory)
	assert.Equal(t, cost.ApplicationCosts[1].CostPercentageByMemory, cost.ApplicationCosts[2].CostPercentageByMemory)
	assert.Equal(t, cost.ApplicationCosts[2].CostPercentageByMemory, cost.ApplicationCosts[3].CostPercentageByMemory)

	assert.Equal(t, 0.25, cost.ApplicationCosts[0].CostPercentageByCPU)
	assert.Equal(t, 0.25, cost.ApplicationCosts[0].CostPercentageByMemory)
}

func Test_cost_one_app_double_requested(t *testing.T) {
	runs := getTestRuns()
	runs[0].Resources[0].Replicas = 4
	runs[1].Resources[0].Replicas = 4
	runs[2].Resources[0].Replicas = 4
	cost := cost_models.NewApplicationCostSet(time.Now().Add(-1), time.Now(), runs, 0, "")

	assert.Equal(t, 0.4, cost.GetCostBy("app-1").CostPercentageByCPU)
	assert.Equal(t, 0.2, cost.GetCostBy("app-2").CostPercentageByCPU)
	assert.Equal(t, 0.2, cost.GetCostBy("app-3").CostPercentageByCPU)
	assert.Equal(t, 0.2, cost.GetCostBy("app-4").CostPercentageByCPU)

	assert.Equal(t, 0.4, cost.GetCostBy("app-1").CostPercentageByMemory)
	assert.Equal(t, 0.2, cost.GetCostBy("app-2").CostPercentageByMemory)
	assert.Equal(t, 0.2, cost.GetCostBy("app-3").CostPercentageByMemory)
	assert.Equal(t, 0.2, cost.GetCostBy("app-4").CostPercentageByMemory)
}

func Test_cost_one_app_no_requested(t *testing.T) {
	runs := getTestRuns()
	runs[0].Resources[0].Replicas = 0
	runs[1].Resources[0].Replicas = 0
	runs[2].Resources[0].Replicas = 0
	cost := cost_models.NewApplicationCostSet(time.Now().Add(-1), time.Now(), runs, 0, "")

	oneThird := float64(1.0) / float64(3.0)
	assert.Equal(t, 0.0, cost.GetCostBy("app-1").CostPercentageByCPU)
	assert.Equal(t, oneThird, cost.GetCostBy("app-2").CostPercentageByCPU)
	assert.Equal(t, oneThird, cost.GetCostBy("app-3").CostPercentageByCPU)
	assert.Equal(t, oneThird, cost.GetCostBy("app-4").CostPercentageByCPU)

	assert.Equal(t, 0.0, cost.GetCostBy("app-1").CostPercentageByMemory)
	assert.Equal(t, oneThird, cost.GetCostBy("app-2").CostPercentageByMemory)
	assert.Equal(t, oneThird, cost.GetCostBy("app-3").CostPercentageByMemory)
	assert.Equal(t, oneThird, cost.GetCostBy("app-4").CostPercentageByMemory)
}

func Test_future_cost_distributed_equally(t *testing.T) {

	run1 := getTestRunForSingleApp("app-1")
	costApp1, _ := cost_models.NewFutureCostEstimate("app-1", run1, subscriptionCost, subscriptionCostCurrency)
	assert.Equal(t, costApp1.Cost, float64(25000))

	run2 := getTestRunForSingleApp("app-2")
	costApp2, _ := cost_models.NewFutureCostEstimate("app-2", run2, subscriptionCost, subscriptionCostCurrency)
	assert.Equal(t, costApp2.Cost, float64(25000))

	run3 := getTestRunForSingleApp("app-3")
	costApp3, _ := cost_models.NewFutureCostEstimate("app-3", run3, subscriptionCost, subscriptionCostCurrency)
	assert.Equal(t, costApp3.Cost, float64(25000))

	run4 := getTestRunForSingleApp("app-4")
	costApp4, _ := cost_models.NewFutureCostEstimate("app-4", run4, subscriptionCost, subscriptionCostCurrency)
	assert.Equal(t, costApp4.Cost, float64(25000))

	// Check that the cost for all applications together covers the total subscription cost
	assert.Equal(t, float64(costApp1.Cost+costApp2.Cost+costApp3.Cost+costApp4.Cost), float64(subscriptionCost))
}

func Test_total_cost_subscription_cost_covered(t *testing.T) {
	runs := getTestRuns()

	cost := cost_models.NewApplicationCostSet(time.Now().Add(-30), time.Now(), runs, subscriptionCost, subscriptionCostCurrency)

	totalCost := float64(0)
	for _, c := range cost.ApplicationCosts {
		totalCost += c.Cost
	}

	assert.Equal(t, totalCost, float64(subscriptionCost))
}

func Test_total_cost_distributed_equally(t *testing.T) {
	runs := getTestRunsWithUnequallyDistributedResources()

	cost := cost_models.NewApplicationCostSet(time.Now().Add(-30), time.Now(), runs, subscriptionCost, subscriptionCostCurrency)

	totalRequestPercentage := float64(0)
	for _, a := range cost.ApplicationCosts {
		parsedPercentage, _ := strconv.ParseFloat(a.Name[6:], 64)
		totalRequestPercentage += parsedPercentage
	}

	for _, appCost := range cost.ApplicationCosts {
		percentageOfTotalCost, _ := strconv.ParseFloat(appCost.Name[6:], 64)
		diffPercentage := percentageOfTotalCost / totalRequestPercentage
		expectedCost := diffPercentage * subscriptionCost
		assert.Equal(t, math.Round(expectedCost), math.Round(appCost.Cost))
	}

}

func getTestRunForSingleApp(appName string) cost_models.Run {
	runs := getTestRuns()

	sort.Slice(runs, func(i, j int) bool {
		return runs[i].ID < runs[j].ID
	})

	return runs[0]
}

func getTestRunsWithUnequallyDistributedResources() []cost_models.Run {
	return []cost_models.Run{
		{
			ID:                    1,
			ClusterCPUMillicore:   20000,
			ClusterMemoryMegaByte: 30000,
			Resources: []cost_models.RequiredResources{
				{
					Application:     "app-1-0.5",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    5000,
					MemoryMegaBytes: 7500,
					Replicas:        2,
				},
				{
					Application:     "app-2-0.2",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    4000,
					MemoryMegaBytes: 6000,
					Replicas:        1,
				},
				{
					Application:     "app-3-0.1",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    1000,
					MemoryMegaBytes: 1500,
					Replicas:        2,
				},
				{
					Application:     "app-4-0.05",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    500,
					MemoryMegaBytes: 750,
					Replicas:        2,
				},
			},
		},
	}
}

func getTestRuns() []cost_models.Run {
	return []cost_models.Run{
		{
			ID:                    1,
			ClusterCPUMillicore:   1000,
			ClusterMemoryMegaByte: 1000,
			Resources: []cost_models.RequiredResources{
				{
					Application:     "app-1",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-2",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-3",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-4",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
			},
		},
		{
			ID:                    2,
			ClusterCPUMillicore:   1000,
			ClusterMemoryMegaByte: 1000,
			Resources: []cost_models.RequiredResources{
				{
					Application:     "app-1",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-2",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-3",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-4",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
			},
		},
		{
			ID:                    3,
			ClusterCPUMillicore:   2000,
			ClusterMemoryMegaByte: 2000,
			Resources: []cost_models.RequiredResources{
				{
					Application:     "app-1",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-2",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-3",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
				{
					Application:     "app-4",
					Environment:     "env-1",
					Component:       "comp-1",
					CPUMillicore:    100,
					MemoryMegaBytes: 100,
					Replicas:        2,
				},
			},
		},
	}
}
