package reportmodels

import (
	"bytes"
	"encoding/csv"
	"testing"

	"github.com/equinor/radix-cost-allocation-api/models"

	"github.com/stretchr/testify/assert"
)

func Test_Created_Report_Exists(t *testing.T) {
	report := NewCostReport(&models.ApplicationCostSet{
		ApplicationCosts: []models.ApplicationCost{
			{Name: "name1"},
			{Name: "name2"},
		},
	})
	fileData := &bytes.Buffer{}

	err := report.Create(fileData)

	assert.Nil(t, err)

	// Add assertion that reads the CSV report and verifies the content
	reader := csv.NewReader(fileData)
	reader.Comma = ';'
	allContent, err := reader.ReadAll()

	assert.Len(t, allContent, 3)

}
