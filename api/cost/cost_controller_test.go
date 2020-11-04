package cost

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	controllertest "github.com/equinor/radix-cost-allocation-api/api/test"
	"github.com/equinor/radix-cost-allocation-api/api/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	appName    = "any-app"
	timeLayout = "2006-01-02"
)

func setupTest() {
	// Set necessary environment variables
	os.Setenv("WHITELIST", "{\"whiteList\": [\"canarycicd-test\",\"canarycicd-test1\",\"canarycicd-test2\",\"canarycicd-test3\",\"radix-api\",\"radix-canary-golang\",\"radix-cost-allocation-api\",\"radix-github-webhook\",\"radix-platform\",\"radix-web-console\"]}")
	os.Setenv("SUBSCRIPTION_COST_VALUE", "100000")
	os.Setenv("SUBSCRIPTION_COST_CURRENCY", "NOK")
}

func TestCostController_Application(t *testing.T) {
	setupTest()

	// Mock setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock auth provider
	fakeAuthProvider := mock.NewMockAuthProvider(ctrl)

	// Create mock idtoken
	fakeIDToken := mock.NewMockIDToken(ctrl)

	fakeIDToken.EXPECT().
		GetClaims(gomock.Any()).
		Return(nil).
		AnyTimes()

	fakeAuthProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(fakeIDToken, nil).
		AnyTimes()

	// Creates a mock Repository
	fakeCostRepo := mock.NewMockCostRepository(ctrl)

	// Generate run with test data
	run := controllertest.ARun().BuildRun()

	// Generate runs with test data
	runs := controllertest.
		AListOfRuns().
		BuildRuns()

	// GetLatestRun() returns a mock run with test data
	fakeCostRepo.EXPECT().
		GetLatestRun().
		Return(*run, nil).
		AnyTimes()

	// GetRunsBetweenTimes() returns mock runs
	fakeCostRepo.EXPECT().
		GetRunsBetweenTimes(gomock.Any(), gomock.Any()).
		Return(runs, nil).
		AnyTimes()

	controllerTestUtils := controllertest.NewTestUtils(NewCostController(fakeCostRepo))
	controllerTestUtils.SetAuthProvider(fakeAuthProvider)

	// Test that futurecost endpoint returns cost for requested application
	t.Run("Futurecost application exists", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/futurecost/%s", appName))
		response := <-responseChannel
		applicationCost := costModels.ApplicationCost{}
		err := controllertest.GetResponseBody(response, &applicationCost)

		assert.Nil(t, err)
		assert.NotNil(t, applicationCost)
		assert.Equal(t, applicationCost.Name, appName)
	})

	// Test that totalcost endpoint returns cost for requested application
	t.Run("Totalcost application exists", func(t *testing.T) {
		from, to := getTimePeriod()
		url := fmt.Sprintf("/api/v1/totalcost/%s?fromTime=%s&toTime=%s", appName, from, to)
		responseChannel := controllerTestUtils.ExecuteRequest("GET", url)
		response := <-responseChannel

		applicationCostSet := costModels.ApplicationCostSet{}
		controllertest.GetResponseBody(response, &applicationCostSet)
		assert.NotNil(t, applicationCostSet)
		applicationCost := applicationCostSet.ApplicationCosts[0]
		assert.Equal(t, applicationCost.Name, appName)
	})

	t.Run("Futurecost estimate is not 0", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/futurecost/%s", appName))
		response := <-responseChannel
		applicationCost := costModels.ApplicationCost{}
		err := controllertest.GetResponseBody(response, &applicationCost)

		assert.Nil(t, err)
		assert.NotNil(t, applicationCost.Cost)
		assert.NotEqual(t, applicationCost.Cost, 0)
	})
}

func TestCostController_Authentication(t *testing.T) {
	setupTest()

	// Mock setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock cost repo
	fakeCostRepo := mock.NewMockCostRepository(ctrl)

	// Create mock auth provider
	fakeAuthProvider := mock.NewMockAuthProvider(ctrl)

	fakeAuthProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(nil, fmt.Errorf("Invalid token")).
		AnyTimes()

	controllerTestUtils := controllertest.NewTestUtils(NewCostController(fakeCostRepo))
	controllerTestUtils.SetAuthProvider(fakeAuthProvider)

	t.Run("Invalid auth header", func(t *testing.T) {
		url := fmt.Sprintf("/api/")
		responseChannel := controllerTestUtils.ExecuteRequest("GET", url)
		response := <-responseChannel

		assert.Equal(t, response.Code, http.StatusUnauthorized)
	})

}

func getTimePeriod() (from, to string) {
	return time.Now().AddDate(0, 0, -30).Format(timeLayout), time.Now().Format(timeLayout)
}
