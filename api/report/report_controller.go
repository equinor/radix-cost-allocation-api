package report

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/equinor/radix-common/models"
	radixhttp "github.com/equinor/radix-common/net/http"
	"github.com/equinor/radix-cost-allocation-api/service"
	log "github.com/sirupsen/logrus"
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
		log.Debugf("Failed to get report. Error: %v", err)
		radixhttp.ErrorResponse(w, r, err)
	}

	radixhttp.ReaderFileResponse(w, &b, fileName, "text/plain; charset=utf-8")
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
