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
// swagger:model CostReport
type CostReport struct {
	// CostReport created date
	//
	// required: true
	PostingDate []string `json:"postingDate"`
	// CostReport document header
	//
	// required: true
	DocumentHeader []string `json:"documentHeader"`
	// CostReport company code
	//
	// required: true
	CompanyCode []string `json:"companyCode"`
	// CostReport WBS
	//
	// required: true
	WBS []string `json:"wbs"`
	// CostReport GLA
	//
	// required: true
	GeneralLedgerAccount []string `json:"generalLedgerAccount"`
	// CostReport cost amount
	//
	// required: true
	Amount []string `json:"amount"`
	// CostReport line text
	//
	// required: true
	LineText []string `json:"lineText"`
}

// NewCostReport constructor
func NewCostReport(appCostSet *costModels.ApplicationCostSet) *CostReport {
	cr := CostReport{}
	for _, appCost := range appCostSet.ApplicationCosts {
		cr.PostingDate = append(cr.PostingDate, time.Now().Format("2006-01-02"))
		cr.Amount = append(cr.Amount, fmt.Sprintf("%.2f", appCost.Cost))
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
	defer writer.Flush()

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

	err = csvWriter.WriteAll(allData)

	if err != nil {
		return err
	}

	return nil

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
