package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/equinor/radix-common/models"
	"github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	"github.com/equinor/radix-cost-allocation-api/router"
	log "github.com/sirupsen/logrus"
)

const reportName = "report.csv"

// Utils Instance variables
type Utils struct {
	controllers  []models.Controller
	authProvider auth.AuthProvider
}

// ReportUtils Instance variables
type ReportUtils struct {
	File *os.File
}

// NewReportTestUtils Constructor
func NewReportTestUtils() ReportUtils {
	file, _ := os.Create(reportName)

	return ReportUtils{
		file,
	}
}

// NewTestUtils Constructor
func NewTestUtils(controllers ...models.Controller) Utils {
	return Utils{
		controllers,
		nil,
	}
}

// SetAuthProvider sets auth provider
func (tu *Utils) SetAuthProvider(ap auth.AuthProvider) {
	tu.authProvider = ap
}

// ExecuteRequest Helper method to issue a http request
func (tu *Utils) ExecuteRequest(method, endpoint string) <-chan *httptest.ResponseRecorder {
	return tu.ExecuteRequestWithParameters(method, endpoint, nil)
}

// ExecuteRequestWithParameters Helper method to issue a http request with payload
func (tu *Utils) ExecuteRequestWithParameters(method, endpoint string, parameters interface{}) <-chan *httptest.ResponseRecorder {
	var reader io.Reader

	if parameters != nil {
		payload, _ := json.Marshal(parameters)
		reader = bytes.NewReader(payload)
	}

	req, _ := http.NewRequest(method, endpoint, reader)
	req.Header.Add("Authorization", "bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyIsImtpZCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyJ9.eyJhdWQiOiIxMjM0NTY3OC0xMjM0LTEyMzQtMTIzNC0xMjM0MjQ1YTJlYzEiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC8xMjM0NTY3OC03NTY1LTIzNDItMjM0Mi0xMjM0MDViNDU5YjAvIiwiaWF0IjoxNTc1MzU1NTA4LCJuYmYiOjE1NzUzNTU1MDgsImV4cCI6MTU3NTM1OTQwOCwiYWNyIjoiMSIsImFpbyI6IjQyYXNkYXMiLCJhbXIiOlsicHdkIl0sImFwcGlkIjoiMTIzNDU2NzgtMTIzNC0xMjM0LTEyMzQtMTIzNDc5MDM5YTkwIiwiYXBwaWRhY3IiOiIwIiwiZmFtaWx5X25hbWUiOiJSYWRpeCIsImdpdmVuX25hbWUiOiJBIFJhZGl4IFVzZXIiLCJoYXNncm91cHMiOiJ0cnVlIiwiaXBhZGRyIjoiMTQzLjk3LjIuMTI5IiwibmFtZSI6IkEgUmFkaXggVXNlciIsIm9pZCI6IjEyMzQ1Njc4LTEyMzQtMTIzNC0xMjM0LTEyMzRmYzhmYTBlYSIsIm9ucHJlbV9zaWQiOiJTLTEtNS0yMS0xMjM0NTY3ODktMTIzNDU2OTc4MC0xMjM0NTY3ODktMTIzNDU2NyIsInNjcCI6InVzZXJfaW1wZXJzb25hdGlvbiIsInN1YiI6IjBoa2JpbEo3MTIzNHpSU3h6eHZiSW1hc2RmZ3N4amI2YXNkZmVOR2FzZGYiLCJ0aWQiOiIxMjM0NTY3OC0xMjM0LTEyMzQtMTIzNC0xMjM0MDViNDU5YjAiLCJ1bmlxdWVfbmFtZSI6IlJBRElYQGVxdWlub3IuY29tIiwidXBuIjoiUkFESVhAZXF1aW5vci5jb20iLCJ1dGkiOiJCUzEyYXNHZHVFeXJlRWNEY3ZoMkFHIiwidmVyIjoiMS4wIn0=.inP8fD7")
	req.Header.Add("Accept", "application/json")

	response := make(chan *httptest.ResponseRecorder)
	go func() {
		rr := httptest.NewRecorder()
		router.NewServer("anyClusterName", tu.authProvider, tu.controllers...).ServeHTTP(rr, req)
		response <- rr
		close(response)
	}()

	return response

}

// GetResponseBody Gets response payload as type
func GetResponseBody(response *httptest.ResponseRecorder, target interface{}) error {
	body, _ := ioutil.ReadAll(response.Body)
	log.Infof(string(body))

	return json.Unmarshal(body, target)
}
