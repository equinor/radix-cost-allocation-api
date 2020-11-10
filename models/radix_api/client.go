package radix_api

import (
	"fmt"
	"os"
	"strings"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

const (
	apiEndpointPatternForContext = "api.%sradix.equinor.com"
	apiEndpointPatternForCluster = "server-radix-api-%s.%s.dev.radix.equinor.com"
)

// RadixAPIClient interface
type RadixAPIClient interface {
	GetRadixApplicationDetails(appParams *application.GetApplicationParams, token string) (*RadixApplicationDetails, error)
	ShowRadixApplications(appParams *platform.ShowApplicationsParams, token string) (*map[string]*RadixApplicationDetails, error)
}

// RadixApplicationDetails instance variables
type RadixApplicationDetails struct {
	Name    string
	Creator string
	Owner   string
	WBS     string
}

// Env instance variables
type Env struct {
	Context        string
	APIEnvironment string
	Cluster        string
}

// Initialize environment variables
func initEnv() *Env {
	var (
		context = os.Getenv("RADIX_CLUSTER_TYPE")
		apiEnv  = os.Getenv("RADIX_ENVIRONMENT")
		cluster = os.Getenv("RADIX_CLUSTER_NAME")
	)

	if context == "" {
		log.Error("'Context' environment variable is not set")
	}

	if apiEnv == "" {
		log.Error("'API-Environment' environment variable is not set")
	}

	if cluster == "" {
		log.Error("'Cluster' environment variables is not set")
	}

	return &Env{
		Context:        context,
		APIEnvironment: apiEnv,
		Cluster:        cluster,
	}

}

// radixAPIClientStruct instance variables
type radixAPIClientStruct struct {
	client *client.Radixapi
	env    *Env
}

// NewRadixAPIClient constructor
func NewRadixAPIClient() RadixAPIClient {
	env := initEnv()
	client := Get(env.Context, env.Cluster, env.APIEnvironment)
	return radixAPIClientStruct{client: client, env: env}
}

func (c radixAPIClientStruct) ShowRadixApplications(appParams *platform.ShowApplicationsParams, token string) (*map[string]*RadixApplicationDetails, error) {
	// Set token in transport
	c.setTransportWithBearerToken(token)

	resp, err := c.client.Platform.ShowApplications(appParams, nil)

	if err != nil || resp == nil {
		return &map[string]*RadixApplicationDetails{}, err
	}

	radixAppMap := make(map[string]*RadixApplicationDetails)
	for _, appSummary := range resp.Payload {
		name := appSummary.Name
		radixAppMap[name] = &RadixApplicationDetails{
			Name: name,
		}
	}

	return &radixAppMap, nil
}

func (c radixAPIClientStruct) GetRadixApplicationDetails(appParams *application.GetApplicationParams, token string) (*RadixApplicationDetails, error) {

	// Set token in transport
	c.setTransportWithBearerToken(token)

	resp, err := c.client.Application.GetApplication(appParams, nil)

	if err != nil || resp == nil {
		return &RadixApplicationDetails{}, err
	}

	appRegistration := resp.Payload.Registration
	return &RadixApplicationDetails{
		Name:    *appRegistration.Name,
		Creator: *appRegistration.Creator,
		Owner:   *appRegistration.Owner,
		WBS:     appRegistration.WBS,
	}, nil
}

// Get Gets API client for current cluster and environment
func Get(context, cluster, environment string) *client.Radixapi {
	var apiEndpoint string

	if cluster != "" {
		apiEndpoint = getAPIEndpointForCluster(cluster, environment)
	} else {
		radixConfig := RadixConfigAccess{}
		startingConfig := radixConfig.GetStartingConfig()

		if strings.TrimSpace(context) == "" {
			context = startingConfig.Config["context"]
		}

		apiEndpoint = getAPIEndpointForContext(context)
	}

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	return client.New(transport, strfmt.Default)
}

func getAPIEndpointForContext(context string) string {
	return fmt.Sprintf(apiEndpointPatternForContext, getPatternForContext(context))
}

func getAPIEndpointForCluster(cluster, environment string) string {
	return fmt.Sprintf(apiEndpointPatternForCluster, environment, cluster)
}

func getPatternForContext(context string) string {
	contextToPattern := make(map[string]string)
	contextToPattern[ContextDevelopment] = "dev."
	contextToPattern[ContextPlayground] = fmt.Sprintf("%s.", ContextPlayground)
	contextToPattern[ContextProdction] = ""
	return contextToPattern[context]
}

func (c *radixAPIClientStruct) setTransportWithBearerToken(token string) {
	var apiEndpoint string

	if c.env.Cluster != "" {
		apiEndpoint = getAPIEndpointForCluster(c.env.Cluster, c.env.APIEnvironment)
	} else {
		radixConfig := RadixConfigAccess{}
		startingConfig := radixConfig.GetStartingConfig()

		if strings.TrimSpace(c.env.Context) == "" {
			c.env.Context = startingConfig.Config["context"]
		}

		apiEndpoint = getAPIEndpointForContext(c.env.Context)
	}

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = httptransport.BearerToken(token)
	c.client.SetTransport(transport)
}
