package cost_models_test

import (
	"github.com/equinor/radix-cost-allocation-api/api/cost/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_cost_all_app_equal_requested(t *testing.T) {
	runs := getTestRuns()
	cost := cost_models.NewApplicationCostSet(time.Now().Add(-1), time.Now(), runs, 0)

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
	cost := cost_models.NewApplicationCostSet(time.Now().Add(-1), time.Now(), runs, 0)

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
	cost := cost_models.NewApplicationCostSet(time.Now().Add(-1), time.Now(), runs, 0)

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
