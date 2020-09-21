package cost

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	"github.com/equinor/radix-cost-allocation-api/models"
	log "github.com/sirupsen/logrus"
)

// CostHandler Instance variables
type Handler struct {
	token string
}

// Init Constructor
func Init(accounts string) Handler {
	return Handler{
		token: accounts,
	}
}

func (costHandler Handler) getToken() string {
	return costHandler.token
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

	applicationCostSet := costModels.NewApplicationCostSet(*fromTime, *toTime, runs, helper.SubscriptionCost, helper.SubscriptionCurrency)
	if appName != nil && !strings.EqualFold(*appName, "") {
		applicationCostSet.ApplicationCosts = costHandler.filterApplicationCostsBy(appName, &applicationCostSet)
	}

	return &applicationCostSet, nil
}

// GetFutureCost estimates cost for the next 30 days based on last run
func (costHandler Handler) GetFutureCost(appName string) (*costModels.ApplicationCost, error) {
	helper := initCostCalculationHelpers()
	defer helper.Client.Close()

	run, err := helper.Client.GetLatestRun(appName)
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

	cost, err := costModels.NewFutureCostEstimate(appName, run, helper.SubscriptionCost, helper.SubscriptionCurrency)

	if err != nil {
		return nil, err
	}

	return &cost, nil
}

func (costHandler Handler) filterApplicationCostsBy(appName *string, cost *costModels.ApplicationCostSet) []costModels.ApplicationCost {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == *appName {
			return []costModels.ApplicationCost{applicationCost}
		}
	}
	return []costModels.ApplicationCost{}
}
