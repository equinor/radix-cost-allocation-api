package cost

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"
	log "github.com/sirupsen/logrus"
)

// Env variables
type Env struct {
	SubscriptionCost     float64
	SubscriptionCurrency string
	Whitelist            *costModels.Whitelist
	Context              string
	APIEnvironment       string
	Cluster              string
}

// Handler Instance variables
type Handler struct {
	repo     models.CostRepository
	env      Env
	accounts models.Accounts
}

// Init Constructor
func Init(repo models.CostRepository, accounts models.Accounts) Handler {
	env := initEnv()
	return Handler{
		repo:     repo,
		env:      *env,
		accounts: accounts,
	}
}

func initEnv() *Env {

	var (
		subCost        = os.Getenv("SUBSCRIPTION_COST_VALUE")
		subCurrency    = os.Getenv("SUBSCRIPTION_COST_CURRENCY")
		whiteList      = os.Getenv("WHITELIST")
		context        = os.Getenv("RADIX_CLUSTER_TYPE")
		apiEnvironment = os.Getenv("RADIX_ENVIRONMENT")
		cluster        = os.Getenv("RADIX_CLUSTER_NAME")
	)

	subscriptionCost, er := strconv.ParseFloat(subCost, 64)
	if er != nil {
		subscriptionCost = 0.0
		log.Info("Subscription Cost is invalid or is not set.")
	}
	if len(subCurrency) == 0 {
		log.Info("Subscription Cost currency is not set.")
	}

	list := &costModels.Whitelist{}
	err := json.Unmarshal([]byte(whiteList), list)

	if err != nil {
		log.Info("Whitelist is not set")
	}

	return &Env{
		SubscriptionCost:     subscriptionCost,
		SubscriptionCurrency: subCurrency,
		Whitelist:            list,
		Context:              context,
		APIEnvironment:       apiEnvironment,
		Cluster:              cluster,
	}
}

func (costHandler *Handler) getToken() string {
	return costHandler.accounts.GetToken()
}

// GetTotalCost handler for GetTotalCost
func (costHandler *Handler) GetTotalCost(fromTime, toTime *time.Time, appName *string) (*costModels.ApplicationCostSet, error) {
	runs, err := costHandler.repo.GetRunsBetweenTimes(fromTime, toTime)
	if err != nil {
		return nil, err
	}

	filteredRuns, err := costHandler.removeWhitelistedAppsFromRun(runs)

	if err != nil {
		return nil, err
	}

	applicationCostSet := costModels.NewApplicationCostSet(*fromTime, *toTime, filteredRuns, costHandler.env.SubscriptionCost, costHandler.env.SubscriptionCurrency)
	if appName != nil && !strings.EqualFold(*appName, "") {
		applicationCostSet.ApplicationCosts = costHandler.filterApplicationCostsBy(appName, &applicationCostSet)
	}

	radixAPIClient := radix_api.GetForToken(costHandler.env.Context, costHandler.env.Cluster, costHandler.env.APIEnvironment, costHandler.getToken())
	rrMap, err := costHandler.getRadixRegistrationMap(radixAPIClient, appName)

	filteredCosts := costHandler.filterApplicationsByAccess(*rrMap, applicationCostSet.ApplicationCosts)
	applicationCostSet.ApplicationCosts = filteredCosts

	err = costHandler.setApplicationProperties(&applicationCostSet.ApplicationCosts, rrMap)

	if err != nil {
		return nil, err
	}

	return &applicationCostSet, nil
}

func (costHandler *Handler) filterApplicationsByAccess(rrMap map[string]*radixApplication, applicationCosts []costModels.ApplicationCost) []costModels.ApplicationCost {
	filteredApplicationCosts := make([]costModels.ApplicationCost, 0)
	for _, applicationCost := range applicationCosts {
		if _, exists := rrMap[applicationCost.Name]; exists {
			filteredApplicationCosts = append(filteredApplicationCosts, applicationCost)
		}
	}

	return filteredApplicationCosts
}

// GetFutureCost estimates cost for the next 30 days based on last run
func (costHandler *Handler) GetFutureCost(appName *string) (*costModels.ApplicationCost, error) {

	run, err := costHandler.repo.GetLatestRun()
	if err != nil {
		log.Info("Could not fetch latest run")
		return nil, errors.New("Failed to fetch resource usage")
	}
	if run.ClusterCPUMillicore == 0 {
		log.Info("Cluster CPU resources are 0")
		return nil, errors.New("Avaliable CPU resources are 0. A cost estimate can not be made")
	}
	if run.ClusterMemoryMegaByte == 0 {
		log.Info("Cluster memory resources are 0")
		return nil, errors.New("Avaliable memory resources are 0. A cost estimate can not be made")
	}

	filteredRun, err := costHandler.removeWhitelistedAppsFromRun([]costModels.Run{run})

	if err != nil {
		return nil, err
	}

	if len(filteredRun) == 0 {
		return nil, fmt.Errorf("Filtering run for application %s returned empty array", *appName)
	}

	run = filteredRun[0]

	cost, err := costModels.NewFutureCostEstimate(*appName, run, costHandler.env.SubscriptionCost, costHandler.env.SubscriptionCurrency)

	if err != nil {
		return nil, err
	}

	radixAPIClient := radix_api.GetForToken(costHandler.env.Context, costHandler.env.Cluster, costHandler.env.APIEnvironment, costHandler.getToken())
	rrMap, err := costHandler.getRadixRegistrationMap(radixAPIClient, appName)

	if err != nil {
		return nil, err
	}

	filteredByAccess := costHandler.filterApplicationsByAccess(*rrMap, []costModels.ApplicationCost{*cost})

	if hasAccessToApp := len(filteredByAccess) > 0; hasAccessToApp {
		return &filteredByAccess[0], nil
	}

	return nil, fmt.Errorf("User does not have access to application %s", *appName)
}

// Whitelist contains list of apps that are not included in cost distribution

func (costHandler *Handler) removeWhitelistedAppsFromRun(runs []costModels.Run) ([]costModels.Run, error) {
	cleanedRuns := runs
	for index, run := range runs {
		cleanedRun := cleanResources(run, costHandler.env.Whitelist)
		cleanedRuns[index].Resources = cleanedRun.Resources
	}

	return cleanedRuns, nil
}

func cleanResources(run costModels.Run, whiteList *costModels.Whitelist) costModels.Run {
	cleanedResources := make([]costModels.RequiredResources, 0)
	for _, resource := range run.Resources {
		if !find(whiteList.List, resource.Application) {
			cleanedResources = append(cleanedResources, resource)
		}
	}
	run.Resources = cleanedResources
	return run
}

func find(list []string, val string) bool {
	for _, item := range list {
		if strings.EqualFold(val, item) {
			return true
		}
	}

	return false
}

func (costHandler *Handler) filterApplicationCostsBy(appName *string, cost *costModels.ApplicationCostSet) []costModels.ApplicationCost {
	for _, applicationCost := range (*cost).ApplicationCosts {
		if applicationCost.Name == *appName {
			return []costModels.ApplicationCost{applicationCost}
		}
	}
	return []costModels.ApplicationCost{}
}

func (costHandler *Handler) setApplicationProperties(applicationCosts *[]costModels.ApplicationCost, rrMap *map[string]*radixApplication) error {
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

func (costHandler *Handler) getRadixRegistrationMap(radixApiClient *client.Radixapi, appName *string) (*map[string]*radixApplication, error) {
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

func (costHandler *Handler) getRadixApplicationDetails(radixApiClient *client.Radixapi, appName *string) (*radixApplication, error) {
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
