package cost_models

import (
	"strings"
	"time"
)

// Run holds all required resources for a time
type Run struct {
	ID                    int64
	MeasuredTimeUTC       time.Time
	ClusterCPUMillicore   int
	ClusterMemoryMegaByte int
	Resources             []RequiredResources
}

// RequiredResources holds required resources for a single component
type RequiredResources struct {
	ID              int64
	WBS             string
	Application     string
	Environment     string
	Component       string
	CPUMillicore    int
	MemoryMegaBytes int
	Replicas        int
}

type Application struct {
	Name                           string
	WBS                            string
	RequestedCPUMillicore          int
	RequestedMemoryMegaByte        int
	RequestedCPUPercentageOfRun    float64
	RequestedMemoryPercentageOfRun float64
}

// CPUWeightInPeriod weight of a run for a period
func (run Run) CPUWeightInPeriod(totalRequestedCPUForPeriod int) float64 {
	if totalRequestedCPUForPeriod == 0 {
		return 1
	}

	return float64(run.ClusterCPUMillicore) / float64(totalRequestedCPUForPeriod)
}

// MemoryWeightInPeriod weight of a run for a period
func (run Run) MemoryWeightInPeriod(totalRequestedMemoryForPeriod int) float64 {
	if totalRequestedMemoryForPeriod == 0 {
		return 1
	}

	return float64(run.ClusterMemoryMegaByte) / float64(totalRequestedMemoryForPeriod)
}

// RequestedCPUByApplications total requested cpu by applications for a run
func (run Run) RequestedCPUByApplications() int {
	cpu := 0
	for _, req := range run.Resources {
		cpu += req.CPUMillicore * req.Replicas
	}
	return cpu
}

// RequestedMemoryByApplications total requested memory by applications for a run
func (run Run) RequestedMemoryByApplications() int {
	memory := 0
	for _, req := range run.Resources {
		memory += req.MemoryMegaBytes * req.Replicas
	}
	return memory
}

// GetApplicationsRequiredResource returns resource requested aggregated on application
func (run Run) GetApplicationsRequiredResource() []Application {
	totalCPURequestedForRun := float64(run.RequestedCPUByApplications())
	totalMemoryRequestedForRun := float64(run.RequestedMemoryByApplications())

	requiredCPU := map[string]int{}
	requiredMemory := map[string]int{}
	wbsCodes := map[string]string{}
	for _, resource := range run.Resources {
		requiredCPU[resource.Application] += resource.CPUMillicore * resource.Replicas
		requiredMemory[resource.Application] += resource.MemoryMegaBytes * resource.Replicas
		wbsCodes[resource.Application] = resource.WBS
	}
	var applications []Application
	for appName, val := range requiredCPU {
		applications = append(applications, Application{
			Name:                           appName,
			WBS:                            wbsCodes[appName],
			RequestedCPUMillicore:          val,
			RequestedMemoryMegaByte:        requiredMemory[appName],
			RequestedCPUPercentageOfRun:    float64(val) / totalCPURequestedForRun,
			RequestedMemoryPercentageOfRun: float64(requiredMemory[appName]) / totalMemoryRequestedForRun,
		})
	}
	return applications
}

// RemoveWhitelistedApplications removes all applications in the whitelist from the run
func (run *Run) RemoveWhitelistedApplications(whiteList *Whitelist) {
	cleanedResources := make([]RequiredResources, 0)
	for _, resource := range run.Resources {
		if !find(whiteList.List, resource.Application) {
			cleanedResources = append(cleanedResources, resource)
		}
	}
	run.Resources = cleanedResources
}

func find(list []string, val string) bool {
	for _, item := range list {
		if strings.EqualFold(val, item) {
			return true
		}
	}

	return false
}
