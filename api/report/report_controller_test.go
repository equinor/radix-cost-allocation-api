package report

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"

	controllerTest "github.com/equinor/radix-cost-allocation-api/api/test"
	mock "github.com/equinor/radix-cost-allocation-api/api/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	appName    = "any-app"
	timeLayout = "2006-01-02"
)

func setupTest(t *testing.T) *controllerTest.Utils {
	os.Setenv("WHITELIST", "{\"whiteList\": [\"canarycicd-test\",\"canarycicd-test1\",\"canarycicd-test2\",\"canarycicd-test3\",\"radix-api\",\"radix-canary-golang\",\"radix-cost-allocation-api\",\"radix-github-webhook\",\"radix-platform\",\"radix-web-console\"]}")
	os.Setenv("SUBSCRIPTION_COST_VALUE", "100000")
	os.Setenv("SUBSCRIPTION_COST_CURRENCY", "NOK")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Creates a mock Repository
	fakeCostRepo := mock.NewMockCostRepository(ctrl)

	// Generate run with test data
	run := controllerTest.ARun().BuildRun()

	// Generate runs with test data
	runs := controllerTest.
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

	controllerTestUtils := controllerTest.NewTestUtils(NewReportController(fakeCostRepo))

	return &controllerTestUtils

}

func TestController_ReportReturned_NotEmpty(t *testing.T) {
	controllerTestUtils := setupTest(t)

	t.Run("CSV report is returned", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/report"))
		response := <-responseChannel

		returnedReport, err := os.Create("response-file.csv")
		defer returnedReport.Close()
		defer os.Remove(returnedReport.Name())
		io.Copy(returnedReport, response.Body)

		openedFile, _ := os.Open(returnedReport.Name())
		reader := csv.NewReader(openedFile)
		allContent, err := reader.ReadAll()

		assert.NotNil(t, allContent)
		assert.Nil(t, err)
		assert.NotNil(t, returnedReport)

	})
}
