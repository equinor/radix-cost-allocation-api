package cost

import (
	"context"
	"fmt"
	"time"

	radixmodels "github.com/equinor/radix-common/models"
	radixhttp "github.com/equinor/radix-common/net/http"
	"github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"
	"github.com/equinor/radix-cost-allocation-api/service"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog"
)

// CostHandler Instance variables
type CostHandler struct {
	accounts    radixmodels.Accounts
	radixapi    radix_api.RadixAPIClient
	costService service.CostService
}

// NewCostHandler Constructor
func NewCostHandler(accounts radixmodels.Accounts, radixapi radix_api.RadixAPIClient, costService service.CostService) CostHandler {
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
func (costHandler *CostHandler) GetTotalCost(ctx context.Context, fromTime, toTime *time.Time, appName *string) (*models.ApplicationCostSet, error) {
	applicationCostSet, err := costHandler.costService.GetCostForPeriod(*fromTime, *toTime)
	if err != nil {
		return nil, err
	}

	if appName != nil {
		applicationCostSet.FilterApplicationCostBy(*appName)
	}

	rrMap, err := costHandler.getRadixRegistrationMap(appName)

	if err != nil {
		zerolog.Ctx(ctx).Info().Err(err).Msg("Could not get application details")
		return nil, err
	}

	filteredCosts := costHandler.filterApplicationsByAccess(rrMap, applicationCostSet.ApplicationCosts)
	applicationCostSet.ApplicationCosts = filteredCosts

	return applicationCostSet, nil
}

// GetFutureCost estimates cost for the next 30 days based on last run
func (costHandler *CostHandler) GetFutureCost(ctx context.Context, appName string) (*models.ApplicationCost, error) {
	cost, err := costHandler.costService.GetFutureCost(appName)
	if err != nil {
		return nil, err
	}

	rrMap, err := costHandler.getRadixRegistrationMap(&appName)

	if err != nil {
		zerolog.Ctx(ctx).Debug().Err(err).Msg("Unable to get application details")
		return nil, err
	}

	filteredByAccess := costHandler.filterApplicationsByAccess(rrMap, []models.ApplicationCost{*cost})

	if hasAccessToApp := len(filteredByAccess) > 0; hasAccessToApp {
		return &filteredByAccess[0], nil
	}

	err = fmt.Errorf("user does not have access to application %s", appName)
	zerolog.Ctx(ctx).Debug().Msg(err.Error())
	return nil, radixhttp.ApplicationNotFoundError("Application was not found.", err)
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

func (costHandler *CostHandler) getRadixRegistrationMap(appName *string) (map[string]*radix_api.RadixApplicationDetails, error) {

	if appName != nil {
		app, err := costHandler.getRadixApplicationDetails(*appName)
		if err != nil {
			return nil, err
		}
		return map[string]*radix_api.RadixApplicationDetails{app.Name: app}, nil
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
