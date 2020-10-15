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
	"github.com/equinor/radix-operator/pkg/apis/defaults"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	clusterName       = "AnyClusterName"
	containerRegistry = "any.container.registry"
	appName           = "any-app"
)

func setupTest(t *testing.T) *controllertest.Utils {
	os.Setenv(defaults.ActiveClusternameEnvironmentVariable, clusterName)
	os.Setenv("WHITELIST", "{\"whiteList\": [\"canarycicd-test\",\"canarycicd-test1\",\"canarycicd-test2\",\"canarycicd-test3\",\"radix-api\",\"radix-canary-golang\",\"radix-cost-allocation-api\",\"radix-github-webhook\",\"radix-platform\",\"radix-web-console\"]}")
	os.Setenv("SUBSCRIPTION_COST_VALUE", "100000")
	os.Setenv("SUBSCRIPTION_COST_CURRENCY", "NOK")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeCostRepo := mock.NewMockRepository(ctrl)

	run := controllertest.ARun().BuildRun()

	fakeCostRepo.EXPECT().
		GetLatestRun().
		Return(*run, nil).
		AnyTimes()

	now := time.Now()
	fakeCostRepo.EXPECT().GetRunsBetweenTimes(now, now.AddDate(0, 0, -30)).Return(nil, nil).AnyTimes()
	fakeCostRepo.EXPECT().CloseDB().Return().AnyTimes()

	fakeRepo := models.CostRepository{Repo: fakeCostRepo}
	test := fakeRepo.Repo

	// controllerTestUtils is used for issuing HTTP request and processing responses
	controllerTestUtils := controllertest.NewTestUtils(NewCostController(&test))

	return &controllerTestUtils
}

func TestGetTotalCost_ApplicationExists(t *testing.T) {
	controllerTestUtils := setupTest(t)

	t.Run("Futurecost application exists", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/futurecost/%s", appName))
		response := <-responseChannel
		applicationCost := costModels.ApplicationCost{}
		err := controllertest.GetResponseBody(response, &applicationCost)

		assert.Nil(t, err)
		assert.NotNil(t, applicationCost)
		assert.Equal(t, applicationCost.Name, appName)
	})

	// Test
	t.Run("Totalcost application exists", func(t *testing.T) {
		fromDate := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
		toDate := time.Now().Format("2006-01-02")
		url := fmt.Sprintf("/api/v1/totalcost/%s?fromTime=%s&toTime=%s", appName, fromDate, toDate)
		responseChannel := controllerTestUtils.ExecuteRequest("GET", url)
		response := <-responseChannel

		applicationCostSet := costModels.ApplicationCostSet{}
		controllertest.GetResponseBody(response, &applicationCostSet)
		assert.NotNil(t, applicationCostSet)
	})
}
