package report

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	reportModels "github.com/equinor/radix-cost-allocation-api/api/report/models"
	"github.com/equinor/radix-cost-allocation-api/api/utils"
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

	var err error

	var subscriptionCost float64
	subscriptionCost, err = strconv.ParseFloat(subCost, 64)
	if err != nil {
		subscriptionCost = 0.0
		log.Info("Subscription Cost is invalid or is not set.")
	}
	if len(subCurrency) == 0 {
		log.Info("Subscription Cost currency is not set.")
	}

	list := &costModels.Whitelist{}
	err = json.Unmarshal([]byte(whiteList), list)

	if err != nil {
		log.Info("Whitelist is not set. ", err)
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
	from, to := utils.GetFirstAndLastOfPreviousMonth()

	runs, err := rh.repo.GetRunsBetweenTimes(from, to)

	if err != nil {
		log.Info("Could not get runs. ", err)
		return err
	}

	cleanedRuns := make([]costModels.Run, 0)
	for _, run := range runs {
		run.RemoveWhitelistedApplications(rh.env.Whitelist)
		cleanedRuns = append(cleanedRuns, run)
	}

	costSet := costModels.NewApplicationCostSet(*from, *to, cleanedRuns, rh.env.SubscriptionCost, rh.env.SubscriptionCurrency)

	report := reportModels.NewCostReport(&costSet)

	return report.Create(out)
}
