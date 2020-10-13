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
	"github.com/equinor/radix-cost-allocation-api/models"
	log "github.com/sirupsen/logrus"
)

// Client interface
type Client interface {
	GetTotalCost(fromTime, toTime *time.Time, appName *string) (*costModels.ApplicationCostSet, error)
	GetFutureCost(appName string) (*costModels.ApplicationCost, error)
}

// Handler Instance variables
type Handler struct {
}

// Init Constructor
func Init() Handler {
	return Handler{}
}

type costCalculationHelper struct {
	Client               *models.SQLClient
	SubscriptionCost     float64
	SubscriptionCurrency string
}

// todo! create write only connection string? dont need read/admin access
const port = 1433

func initCostCalculationHelpers() costCalculationHelper {
	var (
		sqlServer   = os.Getenv("SQL_SERVER")
		sqlDatabase = os.Getenv("SQL_DATABASE")
		sqlUser     = os.Getenv("SQL_USER")
		sqlPassword = os.Getenv("SQL_PASSWORD")
	)
	sqlClient := models.NewSQLClient(sqlServer, sqlDatabase, port, sqlUser, sqlPassword)

	var (
		subscriptionCostEnv         = os.Getenv("SUBSCRIPTION_COST_VALUE")
		subscriptionCostCurrencyEnv = os.Getenv("SUBSCRIPTION_COST_CURRENCY")
	)
	subscriptionCost, er := strconv.ParseFloat(subscriptionCostEnv, 64)
	if er != nil {
		subscriptionCost = 0.0
		log.Info("Subscription Cost is invalid or is not set.")
	}
	if len(subscriptionCostCurrencyEnv) == 0 {
		log.Info("Subscription Cost currency is not set.")
	}

	return costCalculationHelper{
		Client:               &sqlClient,
		SubscriptionCost:     subscriptionCost,
		SubscriptionCurrency: subscriptionCostCurrencyEnv,
	}
}

// GetTotalCost handler for GetTotalCost
func (costHandler Handler) GetTotalCost(fromTime, toTime *time.Time, appName *string) (*costModels.ApplicationCostSet, error) {
	helper := initCostCalculationHelpers()
	defer helper.Client.Close()

	runs, err := helper.Client.GetRunsBetweenTimes(fromTime, toTime)
	if err != nil {
		return nil, err
	}

	filteredRuns, err := costHandler.removeWhitelistedAppsFromRun(runs)

	if err != nil {
		return nil, err
	}

	applicationCostSet := costModels.NewApplicationCostSet(*fromTime, *toTime, filteredRuns, helper.SubscriptionCost, helper.SubscriptionCurrency)
	if appName != nil && !strings.EqualFold(*appName, "") {
		applicationCostSet.ApplicationCosts = costHandler.filterApplicationCostsBy(appName, &applicationCostSet)
	}

	return &applicationCostSet, nil
}

// GetFutureCost estimates cost for the next 30 days based on last run
func (costHandler Handler) GetFutureCost(appName string) (*costModels.ApplicationCost, error) {
	helper := initCostCalculationHelpers()
	defer helper.Client.Close()

	run, err := helper.Client.GetLatestRun()
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

	cost, err := costModels.NewFutureCostEstimate(appName, run, helper.SubscriptionCost, helper.SubscriptionCurrency)

	if err != nil {
		return nil, err
	}

	return cost, nil
}

// Whitelist contains list of apps that are not included in cost distribution

func (costHandler Handler) removeWhitelistedAppsFromRun(runs []costModels.Run) ([]costModels.Run, error) {
	whiteList := os.Getenv("WHITELIST")
	cleanedRuns := runs

	list := &costModels.Whitelist{}
	err := json.Unmarshal([]byte(whiteList), list)

	if err != nil {
		return nil, err
	}

	for index, run := range runs {
		for _, whiteListedApp := range list.List {
			cleanedRun := cleanResources(run, whiteListedApp)
			cleanedRuns[index].Resources = cleanedRun.Resources
		}
	}

	return cleanedRuns, nil
}

func cleanResources(run costModels.Run, app string) costModels.Run {

	for index, resource := range run.Resources {
		if strings.EqualFold(resource.Application, app) {
			run.Resources = remove(run.Resources, index)
			return cleanResources(run, app)
		}
	}

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

func (costHandler Handler) filterApplicationCostsBy(appName *string, cost *costModels.ApplicationCostSet) []costModels.ApplicationCost {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == *appName {
			return []costModels.ApplicationCost{applicationCost}
		}
	}
	return []costModels.ApplicationCost{}
}
