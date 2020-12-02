package radix_api

import (
	"github.com/equinor/radix-cost-allocation-api/models"
	apiClient "github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
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

// radixAPIClientStruct instance variables
type radixAPIClientStruct struct {
	client *apiClient.Radixapi
	env    *models.Env
}

// NewRadixAPIClient constructor
func NewRadixAPIClient(env *models.Env) RadixAPIClient {
	transport := getRadixApiTransport(env)
	return radixAPIClientStruct{
		client: apiClient.New(transport, strfmt.Default),
		env:    env,
	}
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

func (c *radixAPIClientStruct) setTransportWithBearerToken(token string) {
	transport := getRadixApiTransport(c.env)
	transport.DefaultAuthentication = httptransport.BearerToken(token)
	c.client.SetTransport(transport)
}

func getRadixApiTransport(env *models.Env) *httptransport.Runtime {
	return httptransport.New(env.GetRadixAPIURL(), "/api/v1", env.GetRadixAPISchemes())
}
