package report

import (
	"net/http"

	models "github.com/equinor/radix-cost-allocation-api/models"
)

const rootPath = ""

type reportController struct {
	*models.DefaultController
	repo *models.Repository
}

// NewApplicationController constructor
func NewReportController(repo *models.Repository) models.Controller {
	return &reportController{repo: repo}
}

func (rc *reportController) GetRoutes() models.Routes {
	routes := models.Routes{
		{
			Path:        rootPath + "/",
			Method:      "GET",
			HandlerFunc: rc.GetCostReport,
		},
	}
	return routes
}

func (rc *reportController) GetCostReport(w http.ResponseWriter, r *http.Request) {
	handler := Init(rc.repo)
}
