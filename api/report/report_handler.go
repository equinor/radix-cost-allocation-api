package report

import (
	"io"
	"time"

	cost "github.com/equinor/radix-cost-allocation-api/api/cost"
	reportModels "github.com/equinor/radix-cost-allocation-api/api/report/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
)

// ReportHandler instance variables
type ReportHandler struct {
	repo        models.CostRepository
	costHandler cost.CostHandler
}

// Init constructor
func Init(repo models.CostRepository) ReportHandler {
	costHandler := cost.Init(repo)
	return ReportHandler{
		repo:        repo,
		costHandler: costHandler,
	}
}

// GetCostReport creates a CostReport
func (rh *ReportHandler) GetCostReport(out io.Writer) error {
	from, to := getPeriod()
	applicationCosts, err := rh.costHandler.GetTotalCost(from, to, nil)

	if err != nil {
		return err
	}

	report := reportModels.NewCostReport()
	report.Aggregate(*applicationCosts)
	if err != nil {
		return err
	}

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
