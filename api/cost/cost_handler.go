package cost

import (
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/equinor/radix-cost-allocation-api/api/utils"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"
	"github.com/equinor/radix-cost-allocation-api/service"
	log "github.com/sirupsen/logrus"
)

// CostHandler Instance variables
type CostHandler struct {
	accounts    models.Accounts
	radixapi    radix_api.RadixAPIClient
	costService service.CostService
}

// NewCostHandler Constructor
func NewCostHandler(accounts models.Accounts, radixapi radix_api.RadixAPIClient, costService service.CostService) CostHandler {
	return CostHandler{
		accounts:    accounts,
		radixapi:    radixapi,
		costService: costService,
	}
}

func (costHandler *CostHandler) getToken() string {
	return costHandler.accounts.GetToken()
}

// GetTotalCost handler for GetTotalCost
func (costHandler *CostHandler) GetTotalCost(fromTime, toTime *time.Time, appName *string) (*models.ApplicationCostSet, error) {
	applicationCostSet, err := costHandler.costService.GetCostForPeriod(*fromTime, *toTime)
	if err != nil {
		return nil, err
	}

	if appName != nil {
		applicationCostSet.FilterApplicationCostBy(*appName)
	}

	rrMap, err := costHandler.getRadixRegistrationMap(appName)

	if err != nil {
		log.Info("Could not get application details. ", err)
		return nil, err
	}

	filteredCosts := costHandler.filterApplicationsByAccess(*rrMap, applicationCostSet.ApplicationCosts)
	applicationCostSet.ApplicationCosts = filteredCosts

	return applicationCostSet, nil
}

// GetFutureCost estimates cost for the next 30 days based on last run
func (costHandler *CostHandler) GetFutureCost(appName string) (*models.ApplicationCost, error) {
	cost, err := costHandler.costService.GetFutureCost(appName)
	if err != nil {
		return nil, err
	}

	rrMap, err := costHandler.getRadixRegistrationMap(&appName)

	if err != nil {
		log.Debugf("Unable to get application details. Error: %v", err)
		return nil, err
	}

	filteredByAccess := costHandler.filterApplicationsByAccess(*rrMap, []models.ApplicationCost{*cost})

	if hasAccessToApp := len(filteredByAccess) > 0; hasAccessToApp {
		return &filteredByAccess[0], nil
	}

	log.Debugf("User does not have access to application '%s'.", appName)
	return nil, utils.ApplicationNotFoundError("Application was not found.", fmt.Errorf("User does not have access to application %s", appName))
}

func (costHandler *CostHandler) filterApplicationCostsBy(appName *string, cost *models.ApplicationCostSet) []models.ApplicationCost {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == *appName {
			return []models.ApplicationCost{applicationCost}
		}
	}
	return []models.ApplicationCost{}
}

func (costHandler *CostHandler) filterApplicationsByAccess(rrMap map[string]*radix_api.RadixApplicationDetails, applicationCosts []models.ApplicationCost) []models.ApplicationCost {
	filteredApplicationCosts := make([]models.ApplicationCost, 0)
	for _, applicationCost := range applicationCosts {
		if _, exists := rrMap[applicationCost.Name]; exists {
			filteredApplicationCosts = append(filteredApplicationCosts, applicationCost)
		}
	}

	return filteredApplicationCosts
}

func (costHandler *CostHandler) getRadixRegistrationMap(appName *string) (*map[string]*radix_api.RadixApplicationDetails, error) {

	if appName != nil {
		app, err := costHandler.getRadixApplicationDetails(*appName)
		if err != nil {
			return nil, err
		}
		return &map[string]*radix_api.RadixApplicationDetails{app.Name: app}, nil
	}

	showApplicationParams := platform.NewShowApplicationsParams()
	apps, err := costHandler.radixapi.ShowRadixApplications(showApplicationParams, costHandler.getToken())

	if err != nil {
		return nil, err
	}

	return apps, err
}

func (costHandler *CostHandler) getRadixApplicationDetails(appName string) (*radix_api.RadixApplicationDetails, error) {
	getApplicationParams := application.NewGetApplicationParams()
	getApplicationParams.SetAppName(appName)
	token := costHandler.getToken()

	appDetails, err := costHandler.radixapi.GetRadixApplicationDetails(getApplicationParams, token)
	if err != nil || appDetails == nil {
		return nil, err
	}
	return &radix_api.RadixApplicationDetails{
		Name:    appDetails.Name,
		Creator: appDetails.Creator,
		Owner:   appDetails.Owner,
		WBS:     appDetails.WBS,
	}, nil
}
