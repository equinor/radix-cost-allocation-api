package report

import (
	"fmt"
	"net/http"
	"os"

	"github.com/equinor/radix-cost-allocation-api/api/utils"
	models "github.com/equinor/radix-cost-allocation-api/models"
)

const rootPath = ""

type reportController struct {
	*models.DefaultController
	repo models.CostRepository
}

// NewReportController constructor
func NewReportController(repo models.CostRepository) models.Controller {
	return &reportController{repo: repo}
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
func (rc *reportController) GetCostReport(accounts models.Accounts, w http.ResponseWriter, r *http.Request) {
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

	handler := Init(rc.repo)
	fromDate, toDate := utils.GetFirstAndLastOfPreviousMonth()
	file, err := os.Create(fmt.Sprintf("%s-%s.csv", fromDate.Format("2006-01-02"), toDate.Format("2006-01-02")))
	defer os.Remove(file.Name())

	if err != nil {
		utils.ErrorResponse(w, r, err)
	}

	err = handler.GetCostReport(file)

	if err != nil {
		utils.ErrorResponse(w, r, err)
	}

	utils.FileResponse(w, r, file)
}
