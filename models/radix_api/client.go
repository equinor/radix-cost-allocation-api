package radix_api

import (
	"fmt"
	"os"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
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
	ClusterName    string
	DNSZone        string
}

// Initialize environment variables
func initEnv() *Env {
	var (
		context     = os.Getenv("RADIX_CLUSTER_TYPE")
		apiEnv      = os.Getenv("RADIX_ENVIRONMENT")
		clusterName = os.Getenv("RADIX_CLUSTERNAME")
		dnsZone     = os.Getenv("RADIX_DNS_ZONE")
	)

	if context == "" {
		log.Error("'Context' environment variable is not set")
	}

	if apiEnv == "" {
		log.Error("'API-Environment' environment variable is not set")
	}

	if clusterName == "" {
		log.Error("'Cluster' environment variables is not set")
	}

	if dnsZone == "" {
		log.Error("'DNS Zone' environment variables is not set")
	}

	return &Env{
		Context:        context,
		APIEnvironment: apiEnv,
		ClusterName:    clusterName,
		DNSZone:        dnsZone,
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
	client := Get(env.Context, env.ClusterName, env.APIEnvironment, env.DNSZone)
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
func Get(context, cluster, environment, dnsZone string) *client.Radixapi {
	apiEndpoint := getAPIEndpoint(environment, cluster, dnsZone)

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	return client.New(transport, strfmt.Default)
}

func (c *radixAPIClientStruct) setTransportWithBearerToken(token string) {
	apiEndpoint := getAPIEndpoint(c.env.APIEnvironment, c.env.ClusterName, c.env.DNSZone)

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = httptransport.BearerToken(token)
	c.client.SetTransport(transport)
}

func getAPIEndpoint(environment, clusterName, dnsZone string) string {
	return fmt.Sprintf("server-radix-api-%s.%s.%s", environment, clusterName, dnsZone)
}
