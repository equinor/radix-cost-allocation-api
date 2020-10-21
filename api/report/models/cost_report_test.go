package reportmodels

import (
	"encoding/csv"
	"os"
	"testing"

	reportUtils "github.com/equinor/radix-cost-allocation-api/api/test"

	"github.com/stretchr/testify/assert"
)

func setup() (*reportUtils.ReportUtils, func()) {
	utils := reportUtils.NewReportTestUtils()

	teardown := func() {
		utils.File.Close()
		os.Remove(utils.File.Name())
	}

	return &utils, teardown
}

func Test_Created_Report_Exists(t *testing.T) {

	utils, teardown := setup()
	defer teardown()

	appCostSet := reportUtils.AnApplicationCostSet().BuildApplicationCostSet()

	report := NewCostReport()
	report.Aggregate(*appCostSet)

	err := report.Create(utils.File)

	assert.Nil(t, err)
	assert.NotNil(t, utils.File)

	// Add assertion that reads the CSV report and verifies the content
	createdReport, _ := os.Open(utils.File.Name())
	reader := csv.NewReader(createdReport)
	allContent, err := reader.ReadAll()

	assert.NotNil(t, allContent)

	for _, cols := range allContent {
		for _, rows := range cols {
			assert.NotEmpty(t, rows)
		}
	}

	assert.Nil(t, err)

}
