package report

import (
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

func (rc *reportController) GetCostReport(accounts models.Accounts, w http.ResponseWriter, r *http.Request) {
	handler := Init(rc.repo)
	file, err := os.Create("cost-report.csv")
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
