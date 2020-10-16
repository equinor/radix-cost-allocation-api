package cost

import (
	"fmt"
	"os"
	"testing"
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	controllertest "github.com/equinor/radix-cost-allocation-api/api/test"
	mock "github.com/equinor/radix-cost-allocation-api/api/test/mock"
	models "github.com/equinor/radix-cost-allocation-api/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	appName    = "any-app"
	timeLayout = "2006-01-02"
)

func setupTest(t *testing.T) *controllertest.Utils {
	// Set necessary environment variables
	os.Setenv("WHITELIST", "{\"whiteList\": [\"canarycicd-test\",\"canarycicd-test1\",\"canarycicd-test2\",\"canarycicd-test3\",\"radix-api\",\"radix-canary-golang\",\"radix-cost-allocation-api\",\"radix-github-webhook\",\"radix-platform\",\"radix-web-console\"]}")
	os.Setenv("SUBSCRIPTION_COST_VALUE", "100000")
	os.Setenv("SUBSCRIPTION_COST_CURRENCY", "NOK")

	// Mock setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Creates a mock Repository
	fakeCostRepo := mock.NewMockRepository(ctrl)

	// Generate run with test data
	run := controllertest.ARun().BuildRun()

	// Generate runs with test data
	runs := controllertest.
		AListOfRuns().
		BuildRuns()

	// Set a cost period
	from, to := getTimePeriod()
	fromParsed, err := time.Parse(timeLayout, from)
	toParsed, err := time.Parse(timeLayout, to)
	assert.Nil(t, err)

	// GetLatestRun() returns a mock run with test data
	fakeCostRepo.EXPECT().
		GetLatestRun().
		Return(*run, nil).
		AnyTimes()

	// GetRunsBetweenTimes() returns mock runs
	fakeCostRepo.EXPECT().
		GetRunsBetweenTimes(&fromParsed, &toParsed).
		Return(runs, nil).
		AnyTimes()

	// CloseDB() returns nothing and is not applicable for controller testing
	fakeCostRepo.EXPECT().CloseDB().Return().AnyTimes()

	// Assign mock Repository to a CostRepository struct
	fakeRepo := models.CostRepository{Repo: fakeCostRepo}

	// controllerTestUtils is used for issuing HTTP request and processing responses
	controllerTestUtils := controllertest.NewTestUtils(NewCostController(&fakeRepo.Repo))

	return &controllerTestUtils
}

func TestCostController_ApplicationExists(t *testing.T) {
	controllerTestUtils := setupTest(t)

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
}

func getTimePeriod() (from, to string) {
	return time.Now().AddDate(0, 0, -30).Format(timeLayout), time.Now().Format(timeLayout)
}
