package cost

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	"github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"
	"os"
	"strings"
	"time"
)

// CostHandler Instance variables
type Handler struct {
	token string
}

// Init Constructor
func Init(accounts string) Handler {
	return Handler{
		token: accounts,
	}
}

func (costHandler Handler) getToken() string {
	return costHandler.token
}

// todo! create write only connection string? dont need read/admin access
const port = 1433

// GetTotalCost handler for GetTotalCost
func (costHandler Handler) GetTotalCost(fromTime, toTime *time.Time, appName *string) (*costModels.Cost, error) {
	var (
		sqlServer   = os.Getenv("SQL_SERVER")
		sqlDatabase = os.Getenv("SQL_DATABASE")
		sqlUser     = os.Getenv("SQL_USER")
		sqlPassword = os.Getenv("SQL_PASSWORD")
	)
	sqlClient := models.NewSQLClient(sqlServer, sqlDatabase, port, sqlUser, sqlPassword)
	defer sqlClient.Close()

	var (
		context        = os.Getenv("RADIX_CLUSTER_TYPE")
		apiEnvironment = os.Getenv("RADIX_ENVIRONMENT")
		cluster        = os.Getenv("RADIX_CLUSTER_NAME")
	)

	runs, err := sqlClient.GetRunsBetweenTimes(fromTime, toTime)
	if err != nil {
		return nil, err
	}

	cost := costModels.NewCost(*fromTime, *toTime, runs)
	if appName != nil && !strings.EqualFold(*appName, "") {
		cost.ApplicationCosts = costHandler.filterApplicationCostsBy(appName, &cost)
	}

	radixApi := radix_api.GetForToken(context, cluster, apiEnvironment, costHandler.getToken())
	err = costHandler.setApplicationProperties(&cost.ApplicationCosts, radixApi, appName)
	if err != nil {
		return nil, err
	}

	return &cost, nil
}

func (costHandler Handler) filterApplicationCostsBy(appName *string, cost *costModels.Cost) []costModels.ApplicationCost {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == *appName {
			return []costModels.ApplicationCost{applicationCost}
		}
	}
	return []costModels.ApplicationCost{}
}

func (costHandler Handler) setApplicationProperties(applicationCosts *[]costModels.ApplicationCost, radixApi *client.Radixapi, appName *string) error {
	rrMap, err := costHandler.getRadixRegistrationMap(radixApi, appName)
	if err != nil {
		return err
	}
	for idx := range *applicationCosts {
		radixApp, rrExists := (*rrMap)[(*applicationCosts)[idx].Name]
		if !rrExists {
			(*applicationCosts)[idx].Comment = fmt.Sprintf("RadixApplication not found by application name %s.", (*applicationCosts)[idx].Name)
			continue
		}
		(*applicationCosts)[idx].Creator = radixApp.Creator
		(*applicationCosts)[idx].Owner = radixApp.Owner
		(*applicationCosts)[idx].WBS = radixApp.WBS
	}
	return nil
}

type radixApplication struct {
	Name    string
	Creator string
	Owner   string
	WBS     string
}

func (costHandler Handler) getRadixRegistrationMap(radixApiClient *client.Radixapi, appName *string) (*map[string]*radixApplication, error) {
	if appName != nil && !strings.EqualFold(*appName, "") {
		app, err := costHandler.getRadixApplicationDetails(radixApiClient, appName)
		if err != nil {
			return nil, err
		}
		return &map[string]*radixApplication{app.Name: app}, nil
	}

	showApplicationParams := platform.NewShowApplicationsParams()
	resp, err := radixApiClient.Platform.ShowApplications(showApplicationParams, nil)
	if err != nil {
		return nil, err
	}

	radixAppMap := make(map[string]*radixApplication)
	for _, appSummary := range resp.Payload {
		name := appSummary.Name
		radixAppMap[name] = &radixApplication{
			Name: name,
		}
	}
	return &radixAppMap, err
}

func (costHandler Handler) getRadixApplicationDetails(radixApiClient *client.Radixapi, appName *string) (*radixApplication, error) {
	getApplicationParams := application.NewGetApplicationParams()
	getApplicationParams.SetAppName(*appName)
	resp, err := radixApiClient.Application.GetApplication(getApplicationParams, nil)
	if err != nil || resp == nil {
		return nil, err
	}
	ar := resp.Payload.Registration
	return &radixApplication{
		Name:    *ar.Name,
		Creator: *ar.Creator,
		Owner:   *ar.Owner,
		WBS:     ar.WBS,
	}, nil
}
