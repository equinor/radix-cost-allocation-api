package reportmodels

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
)

const (
	companyCode          = "1200"
	generalLedgerAccount = "6541001"
	documentHeader       = "Omnia Radix"
)

// CostReport contains information to be exported to a CSV report
type CostReport struct {
	postingDate          []string
	documentHeader       []string
	companyCode          []string
	wbs                  []string
	generalLedgerAccount []string
	amount               []string
	lineText             []string
}

// NewCostReport constructor
func NewCostReport() *CostReport {
	return &CostReport{}
}

// Create takes the CostReport object and creates a CSV report according to specification
func (cr *CostReport) Create(out io.Writer) error {
	columns := []string{"Posting_Date", "Document_Header", "Company_Code", "WBS", "General_Ledger_Account", "Amount", "Line_Text"}

	writer := bufio.NewWriter(out)
	csvWriter := csv.NewWriter(writer)
	defer writer.Flush()

	err := csvWriter.Write(columns)

	if err != nil {
		return err
	}

	allData := cr.organiseData(len(cr.wbs),
		cr.postingDate,
		cr.documentHeader,
		cr.companyCode,
		cr.wbs,
		cr.generalLedgerAccount,
		cr.amount,
		cr.lineText)

	err = csvWriter.WriteAll(allData)

	if err != nil {
		return err
	}

	return nil

}

// Aggregate ApplicationCostSet data and construct a CostReport
func (cr *CostReport) Aggregate(appCostSet costModels.ApplicationCostSet) {
	for _, appCost := range appCostSet.ApplicationCosts {
		cr.postingDate = append(cr.postingDate, time.Now().Format("2006-01-02"))
		cr.amount = append(cr.amount, fmt.Sprintf("%f", appCost.Cost))
		cr.companyCode = append(cr.companyCode, companyCode)
		cr.documentHeader = append(cr.documentHeader, documentHeader)
		cr.generalLedgerAccount = append(cr.generalLedgerAccount, generalLedgerAccount)
		cr.lineText = append(cr.lineText, appCost.Name)
		cr.wbs = append(cr.wbs, appCost.WBS)
	}
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
