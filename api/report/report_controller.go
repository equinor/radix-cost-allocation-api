package report

import (
	"net/http"

	models "github.com/equinor/radix-cost-allocation-api/models"
)

const rootPath = ""

type reportController struct {
	*models.DefaultController
}

// NewApplicationController constructor
func NewApplicationController() models.Controller {
	return &reportController{}
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

func (rc *reportController) GetCostReport(accounts models.Accounts, w http.ResponseWriter, r *http.Request) {
	handler := Init(accounts.GetToken())

	err := handler.GetCostReport()
}
