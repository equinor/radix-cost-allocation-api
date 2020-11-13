package report

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	reportModels "github.com/equinor/radix-cost-allocation-api/api/report/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
	log "github.com/sirupsen/logrus"
)

// ReportHandler instance variables
type ReportHandler struct {
	repo models.CostRepository
	env  Env
}

// Env variables
type Env struct {
	SubscriptionCost     float64
	SubscriptionCurrency string
	Whitelist            *costModels.Whitelist
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

// Init constructor
func Init(repo models.CostRepository) ReportHandler {
	env := initEnv()
	return ReportHandler{
		repo: repo,
		env:  *env,
	}
}

// GetCostReport creates a CostReport
func (rh *ReportHandler) GetCostReport(out io.Writer) error {
	from, to := getPeriod()

	runs, err := rh.repo.GetRunsBetweenTimes(from, to)

	if err != nil {
		return err
	}

	cleanedRuns := make([]costModels.Run, 0)
	for _, run := range runs {
		run.RemoveWhitelistedApplications(rh.env.Whitelist)
		cleanedRuns = append(cleanedRuns, run)
	}

	costSet := costModels.NewApplicationCostSet(*from, *to, cleanedRuns, rh.env.SubscriptionCost, rh.env.SubscriptionCurrency)

	report := reportModels.NewCostReport(&costSet)

	err = report.Create(out)

	if err != nil {
		return err
	}

	return nil

}

func getPeriod() (*time.Time, *time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	firstOfLastMonth := firstOfMonth.AddDate(0, -1, 0)
	lastOfLastMonth := firstOfLastMonth.AddDate(0, 1, -1)

	return &firstOfLastMonth, &lastOfLastMonth
}
