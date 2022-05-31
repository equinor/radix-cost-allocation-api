package cost

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/equinor/radix-cost-allocation-api/models"
	serviceMock "github.com/equinor/radix-cost-allocation-api/service/mock"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/application"
	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/client/platform"

	controllertest "github.com/equinor/radix-cost-allocation-api/api/test"
	"github.com/equinor/radix-cost-allocation-api/api/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

const (
	appName                      = "any-app"
	timeLayout                   = "2006-01-02"
	applicationIDontHaveAccessTo = "other-app"
)

func Test_ControllerTestSuite(t *testing.T) {
	suite.Run(t, new(controllerTestSuite))
}

type controllerTestSuite struct {
	suite.Suite
	env            *models.Env
	authProvider   *mock.MockAuthProvider
	idToken        *mock.MockIDToken
	radixAPIClient *mock.MockRadixAPIClient
	costService    *serviceMock.MockCostService
}

func (s *controllerTestSuite) SetupTest() {
	os.Setenv("WHITELIST", "{\"whiteList\": [\"canarycicd-test\",\"canarycicd-test1\",\"canarycicd-test2\",\"canarycicd-test3\",\"canarycicd-test4\",\"radix-api\",\"radix-canary-golang\",\"radix-cost-allocation-api\",\"radix-github-webhook\",\"radix-platform\",\"radix-web-console\"]}")
	s.env = models.NewEnv()
	ctrl := gomock.NewController(s.T())
	s.authProvider = mock.NewMockAuthProvider(ctrl)
	s.idToken = mock.NewMockIDToken(ctrl)
	s.radixAPIClient = mock.NewMockRadixAPIClient(ctrl)
	s.costService = serviceMock.NewMockCostService(ctrl)
}

func (s *controllerTestSuite) TearDownTest() {
	os.Clearenv()
}

func (s *controllerTestSuite) Test_FutureCost_ApplicationExist() {
	expected := models.ApplicationCost{Name: appName, WBS: "anywbs"}

	s.authProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(s.idToken, nil).
		Times(1)
	s.costService.EXPECT().
		GetFutureCost(appName).
		Return(&expected, nil).
		Times(1)
	s.radixAPIClient.EXPECT().
		GetRadixApplicationDetails(application.NewGetApplicationParams().WithAppName(appName), gomock.Any()).
		Return(&radix_api.RadixApplicationDetails{Name: appName}, nil).
		Times(1)

	controllerTestUtils := controllertest.NewTestUtils(NewCostController(s.radixAPIClient, s.costService))
	controllerTestUtils.SetAuthProvider(s.authProvider)

	// Test that futurecost endpoint returns cost for requested application
	response := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/futurecost/%s", appName))
	actual := models.ApplicationCost{}
	err := controllertest.GetResponseBody(response, &actual)

	s.Nil(err)
	s.NotNil(actual)
	s.Equal(expected, actual)
}

func (s *controllerTestSuite) Test_TotalCost_ApplicationExist() {
	from, to := getTimePeriod()
	expected := &models.ApplicationCostSet{
		ApplicationCosts: []models.ApplicationCost{{Name: appName, WBS: "anywbs"}},
	}
	s.authProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(s.idToken, nil).
		Times(1)
	s.costService.EXPECT().
		GetCostForPeriod(from, to).
		Return(expected, nil).
		Times(1)
	s.radixAPIClient.EXPECT().
		GetRadixApplicationDetails(application.NewGetApplicationParams().WithAppName(appName), gomock.Any()).
		Return(&radix_api.RadixApplicationDetails{Name: appName}, nil).
		Times(1)

	controllerTestUtils := controllertest.NewTestUtils(NewCostController(s.radixAPIClient, s.costService))
	controllerTestUtils.SetAuthProvider(s.authProvider)

	url := fmt.Sprintf("/api/v1/totalcost/%s?fromTime=%s&toTime=%s", appName, from.Format(timeLayout), to.Format(timeLayout))
	response := controllerTestUtils.ExecuteRequest("GET", url)

	applicationCostSet := models.ApplicationCostSet{}
	controllertest.GetResponseBody(response, &applicationCostSet)
	s.NotNil(applicationCostSet)
	s.Equal(expected.ApplicationCosts[0], applicationCostSet.ApplicationCosts[0])
}

func (s *controllerTestSuite) Test_TotalCost_OnlyAppsWithAccess() {
	from, to := getTimePeriod()
	expected := &models.ApplicationCostSet{
		ApplicationCosts: []models.ApplicationCost{{Name: appName}, {Name: applicationIDontHaveAccessTo}},
	}

	s.authProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(s.idToken, nil).
		Times(1)
	s.costService.EXPECT().
		GetCostForPeriod(from, to).
		Return(expected, nil).
		Times(1)
	s.radixAPIClient.EXPECT().
		ShowRadixApplications(platform.NewShowApplicationsParams(), gomock.Any()).
		Return(map[string]*radix_api.RadixApplicationDetails{appName: {Name: appName}}, nil).
		Times(1)

	controllerTestUtils := controllertest.NewTestUtils(NewCostController(s.radixAPIClient, s.costService))
	controllerTestUtils.SetAuthProvider(s.authProvider)

	response := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/totalcosts?fromTime=%s&toTime=%s", from.Format(timeLayout), to.Format(timeLayout)))

	applicationCostSet := models.ApplicationCostSet{}
	err := controllertest.GetResponseBody(response, &applicationCostSet)
	s.Nil(err)
	s.Equal([]models.ApplicationCost{expected.ApplicationCosts[0]}, applicationCostSet.ApplicationCosts)
}

func (s *controllerTestSuite) Test_Unauthorized() {
	s.authProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(nil, fmt.Errorf("Invalid token")).
		AnyTimes()

	controllerTestUtils := controllertest.NewTestUtils(NewCostController(s.radixAPIClient, s.costService))
	controllerTestUtils.SetAuthProvider(s.authProvider)

	url := fmt.Sprintf("/api/")
	response := controllerTestUtils.ExecuteRequest("GET", url)

	s.Equal(response.Code, http.StatusUnauthorized)

}

func getTimePeriod() (from, to time.Time) {

	f, t := time.Now().AddDate(0, 0, -30), time.Now()
	return f.Truncate(24 * time.Hour).UTC(), t.Truncate(24 * time.Hour).UTC()
}
