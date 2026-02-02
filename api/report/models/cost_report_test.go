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
	allContent, _ := reader.ReadAll()
	assert.Len(t, allContent, 3)

}

func Test_Created_Report_Replaces_Period_With_Comma(t *testing.T) {
	report := NewCostReport(&models.ApplicationCostSet{
		ApplicationCosts: []models.ApplicationCost{
			{Name: "name1", Cost: 1234.56},
			{Name: "name2", Cost: 7890.12},
		},
	})
	fileData := &bytes.Buffer{}

	err := report.Create(fileData)

	assert.Nil(t, err)

	// Add assertion that reads the CSV report and verifies the content
	reader := csv.NewReader(fileData)
	reader.Comma = ';'
	allContent, _ := reader.ReadAll()
	assert.Len(t, allContent, 3)
	assert.Equal(t, "1234,56", allContent[1][5])
	assert.Equal(t, "7890,12", allContent[2][5])

}
