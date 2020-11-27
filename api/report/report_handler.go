package report

import (
	"io"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	reportModels "github.com/equinor/radix-cost-allocation-api/api/report/models"
	"github.com/equinor/radix-cost-allocation-api/api/utils"
	models "github.com/equinor/radix-cost-allocation-api/models"
	log "github.com/sirupsen/logrus"
)

// ReportHandler instance variables
type ReportHandler struct {
	repo models.CostRepository
	env  *models.Env
}

// NewReportHandler constructor
func NewReportHandler(repo models.CostRepository, env *models.Env) *ReportHandler {
	return &ReportHandler{
		repo: repo,
		env:  env,
	}
}

// GetCostReport creates a CostReport
func (rh *ReportHandler) GetCostReport(out io.Writer) error {
	from, to := utils.GetFirstAndLastOfPreviousMonth()

	runs, err := rh.repo.GetRunsBetweenTimes(from, to, nil)

	if err != nil {
		log.Debugf("Could not get runs. Error: %v", err)
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
