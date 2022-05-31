package report

import (
	"bytes"
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
	serviceMock "github.com/equinor/radix-cost-allocation-api/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

const (
	appName         = "any-app"
	timeLayout      = "2006-01-02"
	notValidADGroup = "NOT-VALID-AD-GROUP"
	validADGroup    = "VALID-AD-GROUP"
)

var env *models.Env

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
	os.Setenv("AD_REPORT_READERS", fmt.Sprintf("{\"groups\": [\"%s\"]}", validADGroup))
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

func (s *controllerTestSuite) TestReportController_UnAuthorizedUser_NoAccess() {

	// Create mock auth provider

	c := auth.Claims{Email: "radix_test@equinor.com", Groups: []string{notValidADGroup}}

	s.idToken.EXPECT().
		GetClaims(gomock.Any()).
		SetArg(0, c).
		Return(nil).
		Times(1)

	s.authProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(s.idToken, nil).
		AnyTimes()

	controllerTestUtils := controllerTest.NewTestUtils(NewReportController(s.costService))
	controllerTestUtils.SetAuthProvider(s.authProvider)

	response := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/report"))

	s.Equal(response.Code, http.StatusForbidden)

}

func (s *controllerTestSuite) TestReportController_AuthorizedUser_CanDownload() {
	c := auth.Claims{Email: "radix_test@equinor.com", Groups: []string{validADGroup}}
	cost := &models.ApplicationCostSet{
		ApplicationCosts: []models.ApplicationCost{
			{Name: "app-1", Cost: 100},
			{Name: "app-2", Cost: 200},
		},
	}

	s.idToken.EXPECT().
		GetClaims(gomock.Any()).
		SetArg(0, c).
		Return(nil).
		Times(1)
	s.authProvider.EXPECT().
		VerifyToken(gomock.Any(), gomock.Any()).
		Return(s.idToken, nil).
		Times(1)
	s.costService.EXPECT().
		GetCostForPeriod(gomock.Any(), gomock.Any()).
		Return(cost, nil).
		Times(1)

	controllerTestUtils := controllerTest.NewTestUtils(NewReportController(s.costService))
	controllerTestUtils.SetAuthProvider(s.authProvider)

	response := controllerTestUtils.ExecuteRequest("GET", fmt.Sprintf("/api/v1/report"))
	s.Equal(response.Code, http.StatusOK)
	returnedReport := bytes.Buffer{}
	io.Copy(&returnedReport, response.Body)
	reader := csv.NewReader(&returnedReport)
	reader.Comma = ';'
	allContent, err := reader.ReadAll()
	s.Require().NoError(err)
	s.Len(allContent, 3)
}
