package report

import (
	"os"
	"time"

	cost "github.com/equinor/radix-cost-allocation-api/api/cost"
	reportModels "github.com/equinor/radix-cost-allocation-api/api/report/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
)

const (
	companyCode          = 1200
	generalLedgerAccount = 6541001
	documentHeader       = "Omnia Radix"
)

// ReportHandler instance variables
type ReportHandler struct {
	repo        *models.Repository
	costHandler cost.CostHandler
}

// Init constructor
func Init(repo *models.Repository) ReportHandler {
	costHandler := cost.Init(repo)
	return ReportHandler{
		repo:        repo,
		costHandler: costHandler,
	}
}

// GetCostReport creates a CostReport
func (rh *ReportHandler) GetCostReport() (*os.File, error) {
	from, to := getPeriod()
	applicationCosts, err := rh.costHandler.GetTotalCost(from, to, nil)

	if err != nil {
		return nil, err
	}

	report := reportModels.NewCostReport()
	reportFile, err := os.Create("report.csv")
	if err != nil {
		return nil, err
	}

	costReport, err := report.Create(reportFile)

	if err != nil {
		return nil, err
	}

	return costReport, nil

}

func getPeriod() (*time.Time, *time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	firstOfLastMonth := firstOfMonth.AddDate(0, -1, +1)
	lastOfLastMonth := firstOfLastMonth.AddDate(0, 1, -1)

	return &firstOfLastMonth, &lastOfLastMonth
}
