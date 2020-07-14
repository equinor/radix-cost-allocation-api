package cost

import (
	"fmt"
	"os"
	"testing"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	controllertest "github.com/equinor/radix-cost-allocation-api/api/test"
	"github.com/equinor/radix-operator/pkg/apis/defaults"
	commontest "github.com/equinor/radix-operator/pkg/apis/test"
	"github.com/equinor/radix-operator/pkg/client/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	kubefake "k8s.io/client-go/kubernetes/fake"
)

const (
	clusterName       = "AnyClusterName"
	containerRegistry = "any.container.registry"
	dnsZone           = "dev.radix.equinor.com"
	appAliasDNSZone   = "app.dev.radix.equinor.com"
)

func setupTest() (*commontest.Utils, *controllertest.Utils, *kubefake.Clientset, *fake.Clientset) {
	// Setup
	kubeclient := kubefake.NewSimpleClientset()
	radixclient := fake.NewSimpleClientset()

	// commonTestUtils is used for creating CRDs
	commonTestUtils := commontest.NewTestUtils(kubeclient, radixclient)
	commonTestUtils.CreateClusterPrerequisites(clusterName, containerRegistry)
	os.Setenv(defaults.ActiveClusternameEnvironmentVariable, clusterName)

	// controllerTestUtils is used for issuing HTTP request and processing responses
	controllerTestUtils := controllertest.NewTestUtils(kubeclient, radixclient, NewApplicationController())

	return &commonTestUtils, &controllerTestUtils, kubeclient, radixclient
}

func TestGetTotalCost_ApplicationExists(t *testing.T) {
	_, controllerTestUtils, _, _ := setupTest()

	// Test
	t.Run("matching repo", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/totalcost/%s", "app-name"))
		response := <-responseChannel

		cost := costModels.Cost{}
		controllertest.GetResponseBody(response, &cost)
		assert.NotNil(t, cost)
	})
}
