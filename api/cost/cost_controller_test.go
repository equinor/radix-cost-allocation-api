package cost

import (
	"fmt"
	"os"
	"testing"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	controllertest "github.com/equinor/radix-cost-allocation-api/api/test"
	"github.com/equinor/radix-operator/pkg/apis/defaults"
	commontest "github.com/equinor/radix-operator/pkg/apis/test"
	"github.com/stretchr/testify/assert"
	kubefake "k8s.io/client-go/kubernetes/fake"
)

const (
	clusterName       = "AnyClusterName"
	containerRegistry = "any.container.registry"
)

func setupTest() *controllertest.Utils {
	kubeclient := kubefake.NewSimpleClientset()
	// commonTestUtils is used for creating CRDs
	commonTestUtils := commontest.NewTestUtils(kubeclient, nil)
	commonTestUtils.CreateClusterPrerequisites(clusterName, containerRegistry)
	os.Setenv(defaults.ActiveClusternameEnvironmentVariable, clusterName)

	// controllerTestUtils is used for issuing HTTP request and processing responses
	controllerTestUtils := controllertest.NewTestUtils(NewApplicationController())

	return &controllerTestUtils
}

func TestGetTotalCost_ApplicationExists(t *testing.T) {
	controllerTestUtils := setupTest()

	// Test
	t.Run("matching repo", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/totalcost/%s", "app-name"))
		response := <-responseChannel

		applicationCostSet := costModels.ApplicationCostSet{}
		controllertest.GetResponseBody(response, &applicationCostSet)
		assert.NotNil(t, applicationCostSet)
	})
}
