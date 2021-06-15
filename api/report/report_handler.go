package report

import (
	"io"
	"time"

	reportModels "github.com/equinor/radix-cost-allocation-api/api/report/models"
	"github.com/equinor/radix-cost-allocation-api/service"
)

// ReportHandler instance variables
type ReportHandler struct {
	costService service.CostService
}

// NewReportHandler constructor
func NewReportHandler(costService service.CostService) *ReportHandler {
	return &ReportHandler{
		costService: costService,
	}
}

// GetCostReport creates a CostReport
func (rh *ReportHandler) GetCostReport(out io.Writer, from, to time.Time) error {

	costSet, err := rh.costService.GetCostForPeriod(from, to)
	if err != nil {
		return err
	}

	report := reportModels.NewCostReport(costSet)
	return report.Create(out)
}
