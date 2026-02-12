package reportmodels

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/equinor/radix-cost-allocation-api/models"
)

const (
	companyCode          = "1200"
	generalLedgerAccount = "6541001"
	documentHeader       = "Omnia Radix"
)

// CostReport contains information to be exported to a CSV report
type CostReport struct {
	PostingDate          []string
	DocumentHeader       []string
	CompanyCode          []string
	WBS                  []string
	GeneralLedgerAccount []string
	Amount               []string
	LineText             []string
}

// NewCostReport constructor
func NewCostReport(appCostSet *models.ApplicationCostSet) *CostReport {
	cr := CostReport{}
	for _, appCost := range appCostSet.ApplicationCosts {
		cr.PostingDate = append(cr.PostingDate, time.Now().Format("2006-01-02"))
		cr.Amount = append(cr.Amount, strings.ReplaceAll(fmt.Sprintf("%.2f", appCost.Cost), ".", ","))
		cr.CompanyCode = append(cr.CompanyCode, companyCode)
		cr.DocumentHeader = append(cr.DocumentHeader, documentHeader)
		cr.GeneralLedgerAccount = append(cr.GeneralLedgerAccount, generalLedgerAccount)
		cr.LineText = append(cr.LineText, appCost.Name)
		cr.WBS = append(cr.WBS, appCost.WBS)
	}

	return &cr
}

// Create takes the CostReport object and creates a CSV report according to specification
func (cr *CostReport) Create(out io.Writer) error {
	columns := []string{"Posting_Date", "Document_Header", "Company_Code", "WBS", "General_Ledger_Account", "Amount", "Line_Text"}

	writer := bufio.NewWriter(out)
	csvWriter := csv.NewWriter(writer)
	// Set field seperator to ;
	csvWriter.Comma = ';'
	defer func() { _ = writer.Flush() }()

	err := csvWriter.Write(columns)

	if err != nil {
		return err
	}

	allData := cr.organiseData(len(cr.WBS),
		cr.PostingDate,
		cr.DocumentHeader,
		cr.CompanyCode,
		cr.WBS,
		cr.GeneralLedgerAccount,
		cr.Amount,
		cr.LineText)

	return csvWriter.WriteAll(allData)
}

func (cr *CostReport) organiseData(numberOfRows int, params ...[]string) [][]string {
	numberOfCols := len(params)
	data := make([][]string, numberOfRows)
	for i := 0; i < numberOfRows; i++ {
		for j := 0; j < numberOfCols; j++ {
			data[i] = append(data[i], params[j][i])
		}
	}

	return data
}
