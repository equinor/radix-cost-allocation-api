package cost

import (
	_ "github.com/denisenkom/go-mssqldb"
	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	"github.com/equinor/radix-cost-allocation-api/models"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
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

// todo! create write only connection string? dont need read/admin access
const port = 1433

// GetTotalCost handler for GetTotalCost
func (costHandler Handler) GetTotalCost(fromTime, toTime *time.Time, appName *string) (*costModels.ApplicationCostSet, error) {
	var (
		sqlServer   = os.Getenv("SQL_SERVER")
		sqlDatabase = os.Getenv("SQL_DATABASE")
		sqlUser     = os.Getenv("SQL_USER")
		sqlPassword = os.Getenv("SQL_PASSWORD")
	)
	sqlClient := models.NewSQLClient(sqlServer, sqlDatabase, port, sqlUser, sqlPassword)
	defer sqlClient.Close()

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
	runs, err := sqlClient.GetRunsBetweenTimes(fromTime, toTime)
	if err != nil {
		return nil, err
	}

	applicationCostSet := costModels.NewApplicationCostSet(*fromTime, *toTime, runs, subscriptionCost, subscriptionCostCurrencyEnv)
	if appName != nil && !strings.EqualFold(*appName, "") {
		applicationCostSet.ApplicationCosts = costHandler.filterApplicationCostsBy(appName, &applicationCostSet)
	}

	return &applicationCostSet, nil
}

func (costHandler Handler) filterApplicationCostsBy(appName *string, cost *costModels.ApplicationCostSet) []costModels.ApplicationCost {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == *appName {
			return []costModels.ApplicationCost{applicationCost}
		}
	}
	return []costModels.ApplicationCost{}
}
