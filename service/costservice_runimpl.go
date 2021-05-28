package service

import (
	"errors"
	"strings"
	"time"

	"github.com/equinor/radix-cost-allocation-api/models"
)

type wbsInfo struct {
	wbs          string
	measuredTime time.Time
}

type runCostService struct {
	repo      models.CostRepository
	whitelist models.Whitelist
	cost      float64
	currency  string
}

func NewRunCostService(repo models.CostRepository, whitelist models.Whitelist, cost float64, currency string) CostService {
	return &runCostService{
		repo:      repo,
		whitelist: whitelist,
		cost:      cost,
		currency:  currency,
	}
}

func (s *runCostService) GetCostForPeriod(from, to time.Time) (*models.ApplicationCostSet, error) {
	runs, err := s.repo.GetRunsBetweenTimes(&from, &to)
	if err != nil {
		return nil, err
	}

	cleanedRuns := make([]models.Run, 0)
	for _, run := range runs {
		run.RemoveWhitelistedApplications(&s.whitelist)
		cleanedRuns = append(cleanedRuns, run)
	}

	applicationCostSet := buildApplicationCostSetFromRuns(from, to, cleanedRuns, s.cost, s.currency)
	return &applicationCostSet, nil

}

func (s *runCostService) GetFutureCost(appName string) (*models.ApplicationCost, error) {
	run, err := s.repo.GetLatestRun()
	if err != nil {
		return nil, errors.New("failed to fetch resource usage")
	}
	if run.ClusterCPUMillicore == 0 {
		return nil, errors.New("available CPU resources are 0. A cost estimate can not be made")
	}
	if run.ClusterMemoryMegaByte == 0 {
		return nil, errors.New("available memory resources are 0. A cost estimate can not be made")
	}

	run.RemoveWhitelistedApplications(&s.whitelist)

	return buildFutureCostEstimateFromRun(appName, run, s.cost, s.currency)
}

// NewApplicationCostSet aggregate cost over a time period for applications
func buildApplicationCostSetFromRuns(from, to time.Time, runs []models.Run, subscriptionCost float64, subscriptionCostCurrency string) models.ApplicationCostSet {
	applicationCosts, totalRequestedCPU, totalRequestedMemory := aggregateCostBetweenDatesOnApplications(runs, subscriptionCost, subscriptionCostCurrency)
	cost := models.ApplicationCostSet{
		From:                 from,
		To:                   to,
		ApplicationCosts:     applicationCosts,
		TotalRequestedCPU:    totalRequestedCPU,
		TotalRequestedMemory: totalRequestedMemory,
	}
	return cost
}

// NewFutureCostEstimate aggregate cost data for the last recorded run
func buildFutureCostEstimateFromRun(appName string, run models.Run, subscriptionCost float64, subscriptionCostCurrency string) (*models.ApplicationCost, error) {
	appCost, err := aggregateCostForSingleRun(run, subscriptionCost, subscriptionCostCurrency, appName)

	if err != nil {
		return nil, err
	}

	appCost.AddWBS(run)
	return appCost, nil
}

// aggregateCostBetweenDatesOnApplications calculates cost for an application
func aggregateCostBetweenDatesOnApplications(runs []models.Run, subscriptionCost float64, subscriptionCostCurrency string) ([]models.ApplicationCost, int, int) {
	totalRequestedCPU := totalRequestedCPU(runs)
	totalRequestedMemory := totalRequestedMemoryMegaBytes(runs)
	cpuPercentages := map[string]float64{}
	memoryPercentage := map[string]float64{}
	wbsCodes := map[string]wbsInfo{}

	for _, run := range runs {
		applications := run.GetApplicationsRequiredResource()
		for _, application := range applications {
			if currentWbs, wbsExist := wbsCodes[application.Name]; !wbsExist || run.MeasuredTimeUTC.After(currentWbs.measuredTime) {
				wbsCodes[application.Name] = wbsInfo{wbs: application.WBS, measuredTime: run.MeasuredTimeUTC}
			}
			cpuPercentages[application.Name] += run.CPUWeightInPeriod(totalRequestedCPU) * application.RequestedCPUPercentageOfRun
			memoryPercentage[application.Name] += run.MemoryWeightInPeriod(totalRequestedMemory) * application.RequestedMemoryPercentageOfRun
		}
	}

	var applications []models.ApplicationCost
	for appName, cpu := range cpuPercentages {
		applications = append(applications, models.ApplicationCost{
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
func aggregateCostForSingleRun(run models.Run, subscriptionCost float64, subscriptionCostCurrency string, appName string) (*models.ApplicationCost, error) {
	var costCoverage float64
	var cpuPercentage, memoryPercentage float64
	costDistribution := map[string]float64{}

	if run.ClusterCPUMillicore <= 0 {
		return nil, errors.New("available CPU resources are 0. A cost estimate can not be made")
	}

	if run.ClusterMemoryMegaByte <= 0 {
		return nil, errors.New("available memory resources are 0. A cost estimate can not be made")
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
		return nil, errors.New("no applications requesting resources")
	}

	// Subscriptioncost is not covered in total by the applications
	if costCoverage < 1 {
		costDistribution = scaleDistribution(costDistribution, costCoverage)
	}

	cost := costDistribution[appName] * subscriptionCost

	appCost := models.ApplicationCost{
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

func totalRequestedMemoryMegaBytes(runs []models.Run) int {
	memory := 0
	for _, run := range runs {
		memory += run.ClusterMemoryMegaByte
	}
	return memory
}

// TotalRequestedCPU total requested cpu for runs between from and to datetime
func totalRequestedCPU(runs []models.Run) int {
	cpu := 0
	for _, run := range runs {
		cpu += run.ClusterCPUMillicore
	}
	return cpu
}
