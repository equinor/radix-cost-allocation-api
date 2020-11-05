package radix_api

import (
	"fmt"
	"strings"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"k8s.io/client-go/rest"
)

const (
	apiEndpointPatternForContext = "api.%sradix.equinor.com"
	apiEndpointPatternForCluster = "server-radix-api-%s.%s.dev.radix.equinor.com"

	// TokenEnvironmentName Name of environment variable to load token from
	TokenEnvironmentName = "APP_SERVICE_ACCOUNT_TOKEN"
)

// Get Gets API client for set context
func Get() *client.Radixapi {
	return GetForContext("")
}

// GetForToken Gets API client with passed token
func GetForToken(context, cluster, environment, token string) *client.Radixapi {
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
	transport.DefaultAuthentication = httptransport.BearerToken(token)
	return client.New(transport, strfmt.Default)
}

// GetForContext Gets API client for set context
func GetForContext(context string) *client.Radixapi {
	radixConfig := RadixConfigAccess{}
	startingConfig := radixConfig.GetStartingConfig()

	if strings.TrimSpace(context) == "" {
		context = startingConfig.Config["context"]
	}

	apiEndpoint := getAPIEndpointForContext(context)
	return getClientForEndpoint(apiEndpoint)
}

// GetForCluster Gets API client for cluster
func GetForCluster(cluster, environment string) *client.Radixapi {
	apiEndpoint := getAPIEndpointForCluster(cluster, environment)
	return getClientForEndpoint(apiEndpoint)
}

func getClientForEndpoint(apiEndpoint string) *client.Radixapi {
	radixConfig := RadixConfigAccess{}
	startingConfig := radixConfig.GetStartingConfig()
	persister := PersisterForRadix(radixConfig)
	provider, _ := rest.GetAuthProvider("", startingConfig, persister)

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	transport.Transport = provider.WrapTransport(transport.Transport)
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
