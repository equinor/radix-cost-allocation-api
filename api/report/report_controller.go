package report

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/equinor/radix-common/models"
	"github.com/equinor/radix-cost-allocation-api/api/internal/utils"
	"github.com/equinor/radix-cost-allocation-api/service"
	"github.com/rs/zerolog/log"
)

const rootPath = ""

type reportController struct {
	*models.DefaultController
	costService service.CostService
}

// NewReportController constructor
func NewReportController(costService service.CostService) models.Controller {
	return &reportController{costService: costService}
}

func (rc *reportController) GetRoutes() models.Routes {
	routes := models.Routes{
		{
			Path:        rootPath + "/report",
			Method:      "GET",
			HandlerFunc: rc.GetCostReport,
		},
	}
	return routes
}

// GetCostReport creates a report for all applications for the previous month
func (rc *reportController) GetCostReport(_ models.Accounts, w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /report report getCostReport
	// ---
	// summary: Get cost-report for all applications for the previous month
	// responses:
	//   "200":
	//     description: "Successfully created report"
	//   "401":
	//     description: "Unauthorized"
	//   "403":
	//     description: "Forbidden"
	//   "404":
	//     description: "Not found"

	handler := NewReportHandler(rc.costService)
	fromDate, toDate := getReportFromAndToDate()
	fileName := fmt.Sprintf("%s-%s.csv", fromDate.Format("2006-01-02"), toDate.Format("2006-01-02"))
	var b bytes.Buffer

	err := handler.GetCostReport(&b, fromDate, toDate)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get report")
		utils.ErrorResponseForServer(w, r, fmt.Errorf("failed to get report"))
	}

	utils.ReaderFileResponse(w, r, &b, fileName, "text/plain; charset=utf-8")
}

// from is the first day of the previous month
// to is the first day of current month
// cost is calculated by including the from date up to, but excluding, the to date
func getReportFromAndToDate() (from time.Time, to time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	from = firstOfMonth.AddDate(0, -1, 0)
	to = from.AddDate(0, 1, 0)

	return
}
