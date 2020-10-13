package reportmodels

import (
	"encoding/csv"
	"os"
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

// Create takes the CostReport object and creates a CSV report according to specification
func (cr *CostReport) Create(file *os.File) (*os.File, error) {
	columns := []string{"Posting_Date", "Document_Header", "Company_Code", "WBS", "General_Ledger_Account", "Amount", "Line_Text"}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(columns)

	if err != nil {
		return nil, err
	}

	allData := cr.organiseData(len(cr.wbs),
		cr.postingDate,
		cr.documentHeader,
		cr.companyCode,
		cr.wbs,
		cr.generalLedgerAccount,
		cr.amount,
		cr.lineText)

	err = writer.WriteAll(allData)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (cr *CostReport) organiseData(numberOfRows int, params ...[]string) [][]string {
	numberOfCols := len(params)
	data := make([][]string, numberOfCols)
	for i := 0; i < numberOfRows; i++ {
		for j := 0; j < numberOfCols; j++ {
			data[i] = append(data[i], params[j][i])
		}
	}

	return data
}
