package cost

import (
	"os"

	controllertest "github.com/equinor/radix-cost-allocation-api/api/test"
	mock "github.com/equinor/radix-cost-allocation-api/api/test"
	"github.com/equinor/radix-operator/pkg/apis/defaults"
	commontest "github.com/equinor/radix-operator/pkg/apis/test"
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

	fakeCostRepo := mock.NewFakeCostRepository()
	fakeRepo := fakeCostRepo.Repo

	// controllerTestUtils is used for issuing HTTP request and processing responses
	controllerTestUtils := controllertest.NewTestUtils(NewCostController(&fakeRepo))

	return &controllerTestUtils
}

// func TestGetTotalCost_ApplicationExists(t *testing.T) {
// 	controllerTestUtils := setupTest()

// 	// Test
// 	t.Run("matching repo", func(t *testing.T) {
// 		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/totalcost/%s", "app-name"))
// 		response := <-responseChannel

// 		applicationCostSet := costModels.ApplicationCostSet{}
// 		controllertest.GetResponseBody(response, &applicationCostSet)
// 		assert.NotNil(t, applicationCostSet)
// 	})
// }
