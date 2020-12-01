package cost_models

import (
	"errors"
	"strings"
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

// Whitelist contains list of apps that should not be part of cost distribution
type Whitelist struct {
	// List is the list of apps
	//
	// required: true
	List []string `json:"whiteList"`
}

type wbsInfo struct {
	wbs       string
	knownTime time.Time
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
func NewFutureCostEstimate(appName string, run Run, subscriptionCost float64, subscriptionCostCurrency string) (*ApplicationCost, error) {
	appCost, err := aggregateCostForSingleRun(run, subscriptionCost, subscriptionCostCurrency, appName)

	if err != nil {
		return nil, err
	}

	appCost.AddWBS(run)
	return appCost, nil
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

// FilterApplicationCostBy filters by app name
func (cost *ApplicationCostSet) FilterApplicationCostBy(appName string) {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == appName {
			cost.ApplicationCosts = []ApplicationCost{applicationCost}
			return
		}
	}
	cost.ApplicationCosts = []ApplicationCost{}
}

// AddWBS set WBS to application cost from the run
func (appCost ApplicationCost) AddWBS(run Run) {
	for _, resource := range run.Resources {
		if resource.Application == appCost.Name {
			appCost.WBS = resource.WBS
		}
	}
}

// aggregateCostBetweenDatesOnApplications calculates cost for an application
func aggregateCostBetweenDatesOnApplications(runs []Run, subscriptionCost float64, subscriptionCostCurrency string) ([]ApplicationCost, int, int) {
	totalRequestedCPU := totalRequestedCPU(runs)
	totalRequestedMemory := totalRequestedMemoryMegaBytes(runs)
	cpuPercentages := map[string]float64{}
	memoryPercentage := map[string]float64{}
	wbsCodes := map[string]wbsInfo{}

	for _, run := range runs {
		applications := run.GetApplicationsRequiredResource()
		for _, application := range applications {
			if currentWbs, wbsExist := wbsCodes[application.Name]; !wbsExist || run.MeasuredTimeUTC.After(currentWbs.knownTime) {
				wbsCodes[application.Name] = wbsInfo{wbs: application.WBS, knownTime: run.MeasuredTimeUTC}
			}
			cpuPercentages[application.Name] += run.CPUWeightInPeriod(totalRequestedCPU) * application.RequestedCPUPercentageOfRun
			memoryPercentage[application.Name] += run.MemoryWeightInPeriod(totalRequestedMemory) * application.RequestedMemoryPercentageOfRun
		}
	}

	var applications []ApplicationCost
	for appName, cpu := range cpuPercentages {
		applications = append(applications, ApplicationCost{
			Name:                   appName,
			WBS:                    wbsCodes[appName].wbs,
			Cost:                   (cpu + memoryPercentage[appName]) / 2 * subscriptionCost,
			Currency:               subscriptionCostCurrency,
			CostPercentageByCPU:    cpu,
			CostPercentageByMemory: memoryPercentage[appName],
		})
	}
	return applications, totalRequestedCPU, totalRequestedMemory
}

// Distributes cost to applications for single run
func aggregateCostForSingleRun(run Run, subscriptionCost float64, subscriptionCostCurrency string, appName string) (*ApplicationCost, error) {
	var costCoverage float64
	var cpuPercentage, memoryPercentage float64
	costDistribution := map[string]float64{}

	if run.ClusterCPUMillicore <= 0 {
		return nil, errors.New("Avaliable CPU resources are 0. A cost estimate can not be made")
	}

	if run.ClusterMemoryMegaByte <= 0 {
		return nil, errors.New("Avaliable memory resources are 0. A cost estimate can not be made")
	}

	for _, applicationResources := range run.Resources {

		cpuFraction := float64(applicationResources.CPUMillicore*applicationResources.Replicas) / float64(run.ClusterCPUMillicore)
		memFraction := float64(applicationResources.MemoryMegaBytes*applicationResources.Replicas) / float64(run.ClusterMemoryMegaByte)

		if strings.EqualFold(applicationResources.Application, appName) {
			cpuPercentage += cpuFraction
			memoryPercentage += memFraction
		}

		combined := (cpuFraction + memFraction) / 2
		costDistribution[applicationResources.Application] += combined
		costCoverage += combined
	}

	if costCoverage == 0 {
		return nil, errors.New("No applications requesting resources")
	}

	// Subscriptioncost is not covered in total by the applications
	if costCoverage < 1 {
		costDistribution = scaleDistribution(costDistribution, costCoverage)
	}

	cost := costDistribution[appName] * subscriptionCost

	appCost := ApplicationCost{
		Cost:                   cost,
		Name:                   appName,
		Currency:               subscriptionCostCurrency,
		CostPercentageByCPU:    cpuPercentage,
		CostPercentageByMemory: memoryPercentage,
	}

	return &appCost, nil

}

// Scales the distributed cost up to 100%
func scaleDistribution(distribution map[string]float64, costCoverage float64) map[string]float64 {
	scaled := map[string]float64{}
	for app, fraction := range distribution {
		scaled[app] = fraction / costCoverage
	}

	return scaled
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
