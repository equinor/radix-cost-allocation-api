package reportmodels

import (
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

func Test_Created_Report(t *testing.T) {

	reportUtils, teardown := setup()
	defer teardown()

	report := CostReport{
		amount:               []string{"2023.21", "5091.2", "26982.1"},
		companyCode:          []string{"1200", "1200", "1200"},
		documentHeader:       []string{"Azure Cloud", "Azure Cloud", "Azure Cloud"},
		wbs:                  []string{"A.NES.AD.305", "A.NES.AD.305", "A.NES.AD.305"},
		generalLedgerAccount: []string{"6541001", "6541001", "6541001"},
		lineText:             []string{"8F4920E7-D1A4-470F-AC60-D14467D00350", "8F4920E7-D1A4-470F-AC60-D14467D00350", "8F4920E7-D1A4-470F-AC60-D14467D00350"},
		postingDate:          []string{"2020-10-10", "2020-10-10", "2020-10-10"},
	}

	rep, err := report.Create(reportUtils.File)

	assert.Nil(t, err)
	assert.NotNil(t, rep)

}
