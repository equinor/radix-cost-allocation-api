package cost

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
	log "github.com/sirupsen/logrus"
)

// Env variables
type Env struct {
	SubscriptionCost     float64
	SubscriptionCurrency string
	Whitelist            *costModels.Whitelist
}

// Handler Instance variables
type Handler struct {
	repo models.CostRepository
	env  Env
}

// Init Constructor
func Init(repo models.CostRepository) Handler {
	env := initEnv()
	return Handler{
		repo: repo,
		env:  *env,
	}
}

func initEnv() *Env {

	var (
		subCost     = os.Getenv("SUBSCRIPTION_COST_VALUE")
		subCurrency = os.Getenv("SUBSCRIPTION_COST_CURRENCY")
		whiteList   = os.Getenv("WHITELIST")
	)

	subscriptionCost, er := strconv.ParseFloat(subCost, 64)
	if er != nil {
		subscriptionCost = 0.0
		log.Info("Subscription Cost is invalid or is not set.")
	}
	if len(subCurrency) == 0 {
		log.Info("Subscription Cost currency is not set.")
	}

	list := &costModels.Whitelist{}
	err := json.Unmarshal([]byte(whiteList), list)

	if err != nil {
		log.Info("Whitelist is not set")
	}

	return &Env{
		SubscriptionCost:     subscriptionCost,
		SubscriptionCurrency: subCurrency,
		Whitelist:            list,
	}
}

// GetTotalCost handler for GetTotalCost
func (costHandler *Handler) GetTotalCost(fromTime, toTime *time.Time, appName *string) (*costModels.ApplicationCostSet, error) {
	runs, err := costHandler.repo.GetRunsBetweenTimes(fromTime, toTime)
	if err != nil {
		return nil, err
	}

	filteredRuns, err := costHandler.removeWhitelistedAppsFromRun(runs)

	if err != nil {
		return nil, err
	}

	applicationCostSet := costModels.NewApplicationCostSet(*fromTime, *toTime, filteredRuns, costHandler.env.SubscriptionCost, costHandler.env.SubscriptionCurrency)
	if appName != nil && !strings.EqualFold(*appName, "") {
		applicationCostSet.ApplicationCosts = costHandler.filterApplicationCostsBy(appName, &applicationCostSet)
	}

	return &applicationCostSet, nil
}

// GetFutureCost estimates cost for the next 30 days based on last run
func (costHandler *Handler) GetFutureCost(appName string) (*costModels.ApplicationCost, error) {

	run, err := costHandler.repo.GetLatestRun()
	if err != nil {
		log.Info("Could not fetch latest run")
		return nil, errors.New("Failed to fetch resource usage")
	}
	if run.ClusterCPUMillicore == 0 {
		log.Info("Cluster CPU resources are 0")
		return nil, errors.New("Avaliable CPU resources are 0. A cost estimate can not be made")
	}
	if run.ClusterMemoryMegaByte == 0 {
		log.Info("Cluster memory resources are 0")
		return nil, errors.New("Avaliable memory resources are 0. A cost estimate can not be made")
	}

	filteredRun, err := costHandler.removeWhitelistedAppsFromRun([]costModels.Run{run})

	if err != nil {
		return nil, err
	}

	if len(filteredRun) == 0 {
		return nil, fmt.Errorf("Filtering run for application %s returned empty array", appName)
	}

	run = filteredRun[0]

	cost, err := costModels.NewFutureCostEstimate(appName, run, costHandler.env.SubscriptionCost, costHandler.env.SubscriptionCurrency)

	if err != nil {
		return nil, err
	}

	return cost, nil
}

// Whitelist contains list of apps that are not included in cost distribution

func (costHandler *Handler) removeWhitelistedAppsFromRun(runs []costModels.Run) ([]costModels.Run, error) {
	cleanedRuns := runs
	for index, run := range runs {
		cleanedRun := cleanResources(run, costHandler.env.Whitelist)
		cleanedRuns[index].Resources = cleanedRun.Resources
	}

	return cleanedRuns, nil
}

func cleanResources(run costModels.Run, whiteList *costModels.Whitelist) costModels.Run {
	cleanedResources := make([]costModels.RequiredResources, 0)
	for _, resource := range run.Resources {
		if !find(whiteList.List, resource.Application) {
			cleanedResources = append(cleanedResources, resource)
		}
	}
	run.Resources = cleanedResources
	return run
}

func find(list []string, val string) bool {
	for _, item := range list {
		if strings.EqualFold(val, item) {
			return true
		}
	}

	return false
}

func remove(s []costModels.RequiredResources, i int) []costModels.RequiredResources {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (costHandler *Handler) filterApplicationCostsBy(appName *string, cost *costModels.ApplicationCostSet) []costModels.ApplicationCost {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == *appName {
			return []costModels.ApplicationCost{applicationCost}
		}
	}
	return []costModels.ApplicationCost{}
}
