package report

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/equinor/radix-cost-allocation-api/models"

	controllerTest "github.com/equinor/radix-cost-allocation-api/api/test"
	mock "github.com/equinor/radix-cost-allocation-api/api/test/mock"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	appName         = "any-app"
	timeLayout      = "2006-01-02"
	notValidADGroup = "NOT-VALID-AD-GROUP"
	validADGroup    = "VALID-AD-GROUP"
)

var env *models.Env

func setupTest() {
	os.Setenv("WHITELIST", "{\"whiteList\": [\"canarycicd-test\",\"canarycicd-test1\",\"canarycicd-test2\",\"canarycicd-test3\",\"radix-api\",\"radix-canary-golang\",\"radix-cost-allocation-api\",\"radix-github-webhook\",\"radix-platform\",\"radix-web-console\"]}")
	os.Setenv("SUBSCRIPTION_COST_VALUE", "100000")
	os.Setenv("SUBSCRIPTION_COST_CURRENCY", "NOK")
	os.Setenv("AD_REPORT_READERS", fmt.Sprintf("{\"groups\": [\"%s\"]}", validADGroup))
	env = models.NewEnv()
}

func TestReportController_UnAuthorizedUser_NoAccess(t *testing.T) {
	setupTest()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock auth provider
	fakeAuthProvider := mock.NewMockAuthProvider(ctrl)

	// Create mock idtoken
	fakeIDToken := mock.NewMockIDToken(ctrl)

	c := auth.Claims{Email: "radix_test@equinor.com", Groups: []string{notValidADGroup}}

	fakeIDToken.EXPECT().
		GetClaims(gomock.Any()).
		SetArg(0, c).
		Return(nil).
		Times(1)

	fakeAuthProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(fakeIDToken, nil).
		AnyTimes()

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
		Times(0)

	controllerTestUtils := controllerTest.NewTestUtils(NewReportController(env, fakeCostRepo))
	controllerTestUtils.SetAuthProvider(fakeAuthProvider)

	t.Run("User without required AD group can't download report", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/report"))
		response := <-responseChannel

		assert.Equal(t, response.Code, http.StatusForbidden)
	})

}

func TestReportController_AuthorizedUser_CanDownload(t *testing.T) {
	setupTest()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock auth provider
	fakeAuthProvider := mock.NewMockAuthProvider(ctrl)

	// Create mock idtoken
	fakeIDToken := mock.NewMockIDToken(ctrl)

	c := auth.Claims{Email: "radix_test@equinor.com", Groups: []string{validADGroup}}

	fakeIDToken.EXPECT().
		GetClaims(gomock.Any()).
		SetArg(0, c).
		Return(nil).
		Times(1)

	fakeAuthProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(fakeIDToken, nil).
		AnyTimes()

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
		Times(1)

	controllerTestUtils := controllerTest.NewTestUtils(NewReportController(env, fakeCostRepo))
	controllerTestUtils.SetAuthProvider(fakeAuthProvider)

	t.Run("User with correct AD group can download report", func(t *testing.T) {
		responseChannel := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/report"))
		response := <-responseChannel

		assert.Equal(t, response.Code, http.StatusOK)

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
