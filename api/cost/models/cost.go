package cost_models

import (
	"time"
)

// ApplicationCostSet details of application cost set
// swagger:model ApplicationCostSet
type ApplicationCostSet struct {

	// ApplicationCostSet period started From
	//
	// required: true
	From time.Time `json:"from"`

	// ApplicationCostSet period continued To
	//
	// required: true
	To time.Time `json:"to"`

	// ApplicationCosts with costs.
	//
	// required: true
	ApplicationCosts []ApplicationCost `json:"applicationCosts"`

	// TotalRequestedCPU within the period.
	//
	// required: true
	TotalRequestedCPU int `json:"totalRequestedCpu"`

	// TotalRequestedMemory within the period.
	//
	// required: true
	TotalRequestedMemory int `json:"totalRequestedMemory"`
}

// ApplicationCost details of one application cost
// swagger:model ApplicationCost
type ApplicationCost struct {
	// Name of the application
	//
	// required: true
	// example: radix-canary-golang
	Name string `json:"name"`
	// Owner of the application (email). Can be a single person or a shared group email.
	//
	// required: false
	Owner string `json:"owner"`

	// Creator of the application.
	//
	// required: false
	Creator string `json:"creator"`

	// WBS for the application.
	//
	// required: false
	WBS string `json:"wbs"`

	// CostPercentageByCPU is cost percentage by CPU for the application.
	//
	// required: true
	CostPercentageByCPU float64 `json:"costPercentageByCpu"`

	// CostPercentageByMemory is cost percentage by memory for the application
	//
	// required: true
	CostPercentageByMemory float64 `json:"costPercentageByMemory"`

	// Comment regarding cost
	//
	// required: false
	Comment string `json:"comment"`

	// Cost
	//
	// required: true
	Cost float64 `json:"cost"`

	// Cost currency
	//
	// required: true
	Currency string `json:"currency"`
}

// NewApplicationCostSet aggregate cost over a time period for applications
func NewApplicationCostSet(from, to time.Time, runs []Run, subscriptionCost float64, subscriptionCostCurrency string) ApplicationCostSet {
	applicationCosts, totalRequestedCPU, totalRequestedMemory := aggregateCostBetweenDatesOnApplications(runs, subscriptionCost, subscriptionCostCurrency)
	cost := ApplicationCostSet{
		From:                 from,
		To:                   to,
		ApplicationCosts:     applicationCosts,
		TotalRequestedCPU:    totalRequestedCPU,
		TotalRequestedMemory: totalRequestedMemory,
	}
	return cost
}

// NewFutureCostEstimate aggregate cost data for the last recorded run
func NewFutureCostEstimate(appName string, run Run, subscriptionCost float64, subscriptionCostCurrency string) ApplicationCost {
	appCost := aggregateCostForSingleRun(run, subscriptionCost, subscriptionCostCurrency, appName)
	return appCost
}

// GetCostBy returns application by appName
func (cost ApplicationCostSet) GetCostBy(appName string) *ApplicationCost {
	for _, app := range cost.ApplicationCosts {
		if app.Name == appName {
			return &app
		}
	}
	return nil
}

// aggregateCostBetweenDatesOnApplications calculates cost for an application
func aggregateCostBetweenDatesOnApplications(runs []Run, subscriptionCost float64, subscriptionCostCurrency string) ([]ApplicationCost, int, int) {
	totalRequestedCPU := totalRequestedCPU(runs)
	totalRequestedMemory := totalRequestedMemoryMegaBytes(runs)
	cpuPercentages := map[string]float64{}
	memoryPercentage := map[string]float64{}
	wbsCodes := map[string]string{}

	for _, run := range runs {
		applications := run.GetApplicationsRequiredResource()
		for _, application := range applications {
			wbsCodes[application.Name] = application.WBS
			cpuPercentages[application.Name] += run.CPUWeightInPeriod(totalRequestedCPU) * application.RequestedCPUPercentageOfRun
			memoryPercentage[application.Name] += run.MemoryWeightInPeriod(totalRequestedMemory) * application.RequestedMemoryPercentageOfRun
		}
	}

	var applications []ApplicationCost
	for appName, cpu := range cpuPercentages {
		applications = append(applications, ApplicationCost{
			Name:                   appName,
			WBS:                    wbsCodes[appName],
			Cost:                   cpu * subscriptionCost,
			Currency:               subscriptionCostCurrency,
			CostPercentageByCPU:    cpu,
			CostPercentageByMemory: memoryPercentage[appName],
		})
	}
	return applications, totalRequestedCPU, totalRequestedMemory
}

func aggregateCostForSingleRun(run Run, subscriptionCost float64, subscriptionCostCurrency string, appName string) ApplicationCost {
	var totalRequestedCPU int
	var totalRequestedMemory int

	for _, applicationResources := range run.Resources {
		if applicationResources.Application != appName {
			continue
		}

		totalRequestedCPU += applicationResources.CPUMillicore * applicationResources.Replicas
		totalRequestedMemory += applicationResources.MemoryMegaBytes * applicationResources.Replicas

	}

	cpuWeight := run.CPUWeightInPeriod(totalRequestedCPU)
	memoryWeight := run.MemoryWeightInPeriod(totalRequestedMemory)

	cpuPercentage := cpuWeight * float64(run.ClusterCPUMillicore)
	memoryPercentage := memoryWeight * float64(run.ClusterMemoryMegaByte)

	costWeight := (cpuPercentage + memoryPercentage) / 2
	cost := costWeight * subscriptionCost

	return ApplicationCost{
		Cost:     cost,
		Name:     appName,
		Currency: subscriptionCostCurrency,
	}

}

func totalRequestedMemoryMegaBytes(runs []Run) int {
	memory := 0
	for _, run := range runs {
		memory += run.ClusterMemoryMegaByte
	}
	return memory
}

// TotalRequestedCPU total requested cpu for runs between from and to datetime
func totalRequestedCPU(runs []Run) int {
	cpu := 0
	for _, run := range runs {
		cpu += run.ClusterCPUMillicore
	}
	return cpu
}
