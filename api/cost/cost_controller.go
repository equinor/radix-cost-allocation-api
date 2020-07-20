package cost

import (
	"github.com/equinor/radix-cost-allocation-api/api/utils"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

const rootPath = ""

type costController struct {
	*models.DefaultController
}

// NewApplicationController Constructor
func NewApplicationController() models.Controller {
	return &costController{}
}

// GetRoutes List the supported routes of this controller
func (costController *costController) GetRoutes() models.Routes {
	routes := models.Routes{
		models.Route{
			Path:        rootPath + "/totalcost/{appName}",
			Method:      "GET",
			HandlerFunc: costController.GetTotalCost,
		},
		models.Route{
			Path:        rootPath + "/totalcosts/",
			Method:      "GET",
			HandlerFunc: costController.GetTotalCosts,
		},
	}

	return routes
}

// GetTotalCost for an application
func (costController *costController) GetTotalCost(accounts models.Accounts, w http.ResponseWriter, r *http.Request) {
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
	//   description: Get cost from fromTime (example 2020-03-18T07:20:41+00:00)
	//   type: string
	//   format: date-time
	//   required: false
	// - name: toTime
	//   in: query
	//   description: Get cost to toTime (example 2020-03-18T07:20:41+00:00)
	//   type: string
	//   format: date-time
	//   required: false
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
	//        "$ref": "#/definitions/Cost"
	//   "401":
	//     description: "Unauthorized"
	//   "404":
	//     description: "Not found"
	appName := mux.Vars(r)["appName"]

	fromTime, toTime, err := getCostPeriod(w, r)
	if err != nil {
		utils.ErrorResponse(w, r, err)
		return
	}

	handler := Init(accounts)
	cost, err := handler.GetTotalCost(appName, fromTime, toTime)

	if err != nil {
		utils.ErrorResponse(w, r, err)
		return
	}

	utils.JSONResponse(w, r, cost)
}

// GetTotalCosts for period
func (costController *costController) GetTotalCosts(accounts models.Accounts, w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /totalcosts/ cost getTotalCosts
	// ---
	// summary: Gets the total cost for an application
	// parameters:
	// - name: fromTime
	//   in: query
	//   description: Get cost from fromTime (example 2020-03-18T07:20:41+00:00)
	//   type: string
	//   format: date-time
	//   required: false
	// - name: toTime
	//   in: query
	//   description: Get cost to toTime (example 2020-03-18T07:20:41+00:00)
	//   type: string
	//   format: date-time
	//   required: false
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
	//        "$ref": "#/definitions/Cost"
	//   "401":
	//     description: "Unauthorized"
	//   "404":
	//     description: "Not found"
	fromTime, toTime, err := getCostPeriod(w, r)
	if err != nil {
		utils.ErrorResponse(w, r, err)
		return
	}

	handler := Init(accounts)
	cost, err := handler.GetTotalCosts(fromTime, toTime)

	if err != nil {
		utils.ErrorResponse(w, r, err)
		return
	}

	utils.JSONResponse(w, r, cost)
}

func getCostPeriod(w http.ResponseWriter, r *http.Request) (*time.Time, *time.Time, error) {
	fromTime, err := getTimeFromRequest(r, "fromTime")
	if err != nil {
		utils.ErrorResponse(w, r, err)
		return nil, nil, err
	}

	toTime, err := getTimeFromRequest(r, "toTime")
	if err != nil {
		utils.ErrorResponse(w, r, err)
		return nil, nil, err
	}

	return fromTime, toTime, nil
}

func getTimeFromRequest(r *http.Request, argName string) (*time.Time, error) {
	timeString := r.FormValue(argName)
	var timeValue time.Time
	if !strings.EqualFold(strings.TrimSpace(timeString), "") {
		var err error
		timeValue, err = utils.ParseTimestamp(timeString)
		if err != nil {
			return nil, err
		}
	}
	return &timeValue, nil
}
