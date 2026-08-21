package cost

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/equinor/radix-cost-allocation-api/internal/accounts"
	"github.com/equinor/radix-cost-allocation-api/internal/controller"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/service"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

const rootPath = ""

type costController struct {
	*controller.DefaultController
	radixapi    radix_api.RadixAPIClient
	costService service.CostService
}

// NewCostController Constructor
func NewCostController(radixapi radix_api.RadixAPIClient, costService service.CostService) controller.Controller {
	return &costController{radixapi: radixapi, costService: costService}
}

// GetRoutes List the supported routes of this controller
func (c *costController) GetRoutes() controller.Routes {
	routes := controller.Routes{
		controller.Route{
			Path:        rootPath + "/totalcosts",
			Method:      "GET",
			HandlerFunc: c.GetTotalCosts,
		},
		controller.Route{
			Path:        rootPath + "/totalcost/{appName}",
			Method:      "GET",
			HandlerFunc: c.GetTotalCost,
		},
		controller.Route{
			Path:        rootPath + "/futurecost/{appName}",
			Method:      "GET",
			HandlerFunc: c.GetFutureCost,
		},
	}

	return routes
}

// GetTotalCosts for all applications for period
func (c *costController) GetTotalCosts(accounts accounts.Accounts, w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /totalcosts cost getTotalCosts
	// ---
	// summary: Gets the total cost for an application
	// parameters:
	// - name: fromTime
	//   in: query
	//   description: Get cost from fromTime (example 2020-03-18 or 2020-03-18T07:20:41+01:00)
	//   type: string
	//   format: date-time
	//   required: true
	// - name: toTime
	//   in: query
	//   description: Get cost to toTime (example 2020-09-18 or 2020-09-18T07:20:41+01:00)
	//   type: string
	//   format: date-time
	//   required: true
	// - name: Impersonate-User
	//   in: header
	//   description: Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	//   type: string
	//   required: false
	// - name: Impersonate-Group
	//   in: header
	//   description: Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	//   type: string
	//   required: false
	// responses:
	//   "200":
	//     description: "Successful get cost"
	//     schema:
	//        "$ref": "#/definitions/ApplicationCostSet"
	//   "401":
	//     description: "Unauthorized"
	//   "404":
	//     description: "Not found"
	c.getTotalCosts(accounts, w, r, nil)
}

// GetTotalCost for an application for period
func (c *costController) GetTotalCost(accounts accounts.Accounts, w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /totalcost/{appName} cost getTotalCost
	// ---
	// summary: Gets the total cost for an application
	// parameters:
	// - name: appName
	//   in: path
	//   description: Name of application
	//   type: string
	//   required: true
	// - name: fromTime
	//   in: query
	//   description: Get cost from fromTime (example 2020-03-18 or 2020-03-18T07:20:41+01:00)
	//   type: string
	//   format: date-time
	//   required: true
	// - name: toTime
	//   in: query
	//   description: Get cost to toTime (example 2020-09-18 or 2020-09-18T07:20:41+01:00)
	//   type: string
	//   format: date-time
	//   required: true
	// - name: Impersonate-User
	//   in: header
	//   description: Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	//   type: string
	//   required: false
	// - name: Impersonate-Group
	//   in: header
	//   description: Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	//   type: string
	//   required: false
	// responses:
	//   "200":
	//     description: "Successful get cost"
	//     schema:
	//        "$ref": "#/definitions/ApplicationCostSet"
	//   "401":
	//     description: "Unauthorized"
	//   "404":
	//     description: "Not found"
	appName := mux.Vars(r)["appName"]
	c.getTotalCosts(accounts, w, r, &appName)
}

func (c *costController) GetFutureCost(accounts accounts.Accounts, w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /futurecost/{appName} cost getFutureCost
	// ---
	// summary: Gets the estimated future cost for an application
	// parameters:
	// - name: appName
	//   in: path
	//   description: Name of application
	//   type: string
	//   required: true
	// - name: Impersonate-User
	//   in: header
	//   description: Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	//   type: string
	//   required: false
	// - name: Impersonate-Group
	//   in: header
	//   description: Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	//   type: string
	//   required: false
	// responses:
	//   "200":
	//     description: "Successful get cost"
	//     schema:
	//        "$ref": "#/definitions/ApplicationCost"
	//   "401":
	//     description: "Unauthorized"
	//   "404":
	//     description: "Not found"
	appName := mux.Vars(r)["appName"]
	c.getFutureCost(accounts, w, r, appName)
}

func (c *costController) getFutureCost(accounts accounts.Accounts, w http.ResponseWriter, r *http.Request, appName string) {
	handler := NewCostHandler(accounts, c.radixapi, c.costService)
	cost, err := handler.GetFutureCost(r.Context(), appName)

	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("failed to get future cost")
		c.ErrorResponseForServer(w, r, fmt.Errorf("failed to get future cost"))
		return
	}

	c.JSONResponse(w, r, &cost)
}

func (c *costController) getTotalCosts(accounts accounts.Accounts, w http.ResponseWriter, r *http.Request, appName *string) {
	fromTime, toTime, err := getCostPeriod(r)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("failed to get total cost period")
		c.ErrorResponseForServer(w, r, fmt.Errorf("failed to get total cost period"))
		return
	}

	handler := NewCostHandler(accounts, c.radixapi, c.costService)
	cost, err := handler.GetTotalCost(r.Context(), fromTime, toTime, appName)
	if err != nil {
		zerolog.Ctx(r.Context()).Error().Err(err).Msg("failed to get total cost")
		c.ErrorResponseForServer(w, r, fmt.Errorf("failed to get total cost"))
		return
	}

	c.JSONResponse(w, r, cost)
}

func getCostPeriod(r *http.Request) (*time.Time, *time.Time, error) {
	fromTime, err := getTimeFromRequest(r, "fromTime")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get cost period. Error: %v", err)
	}

	toTime, err := getTimeFromRequest(r, "toTime")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get cost period time to. Error: %v", err)
	}

	return fromTime, toTime, nil
}

func getTimeFromRequest(r *http.Request, argName string) (*time.Time, error) {
	timeString := r.FormValue(argName)
	var timeValue time.Time
	if strings.EqualFold(strings.TrimSpace(timeString), "") {
		return nil, fmt.Errorf("missed argument %s", argName)
	}
	var err error
	if len(timeString) == 10 {
		timeValue, err = time.Parse("2006-01-02", timeString)
	} else {
		timeValue, err = time.Parse(time.RFC3339, timeString)
	}
	if err != nil {
		return nil, err
	}
	return &timeValue, nil
}
