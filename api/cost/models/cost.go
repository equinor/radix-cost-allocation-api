package cost_models

import "time"

// Cost details of cost
// swagger:model Cost
type Cost struct {

	// Cost period started From
	//
	// required: true
	From time.Time

	// Cost period continued To
	//
	// required: true
	To time.Time

	// ApplicationCosts with costs.
	//
	// required: true
	ApplicationCosts []ApplicationCost

	// Runs of ApplicationCosts.
	//
	// required: true
	runs []Run
}

type ApplicationCost struct {
	// Name of the application
	//
	// required: true
	// example: radix-canary-golang
	Name string
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
	WBS string

	// CostPercentageByCPU is cost percentage by CPU for the application.
	//
	// required: true
	CostPercentageByCPU float64

	// CostPercentageByMemory is cost percentage by memory for the application
	//
	// required: true
	CostPercentageByMemory float64

	// Comment regarding cost
	//
	// required: false
	Comment string
}

// NewCost aggregate cost over a time period for applications
func NewCost(from, to time.Time, runs []Run) Cost {
	cost := Cost{
		From:             from,
		To:               to,
		ApplicationCosts: aggregateCostBetweenDatesOnApplications(runs),
		runs:             runs,
	}
	return cost
}

// GetCostBy returns application by appName
func (cost Cost) GetCostBy(appName string) *ApplicationCost {
	for _, app := range cost.ApplicationCosts {
		if app.Name == appName {
			return &app
		}
	}
	return nil
}

// aggregateCostBetweenDatesOnApplications calculates cost for an application
func aggregateCostBetweenDatesOnApplications(runs []Run) []ApplicationCost {
	totalRequestedCPU := totalRequestedCPU(runs)
	totalRequestedMemory := totalRequestedMemoryMegaBytes(runs)
	cpuPercentages := map[string]float64{}
	memoryPercentage := map[string]float64{}

	for _, runs := range runs {
		applications := runs.GetApplicationsRequiredResource()
		for _, application := range applications {
			cpuPercentages[application.Name] += runs.CPUWeightInPeriod(totalRequestedCPU) * application.RequestedCPUPercentageOfRun
			memoryPercentage[application.Name] += runs.MemoryWeightInPeriod(totalRequestedMemory) * application.RequestedMemoryPercentageOfRun
		}
	}

	var applications []ApplicationCost
	for appName, cpu := range cpuPercentages {
		applications = append(applications, ApplicationCost{
			Name: appName,
			//TODO: add owner, creator
			CostPercentageByCPU:    cpu,
			CostPercentageByMemory: memoryPercentage[appName],
		})
	}
	return applications
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
