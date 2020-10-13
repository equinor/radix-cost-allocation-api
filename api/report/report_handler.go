package report

import (
	"time"

	cost_handler "github.com/equinor/radix-cost-allocation-api/api/cost"
)

type CostClient interface {
	GetCost() float64
}

type costClient struct {
	sqlClient string
}

type mockCostClient struct {
	data []float64
}

func NewMockCostClient(data []float64) mockCostClient {
	return mockCostClient{data}
}

func NewCostClient(sqlClient string) costClient {
	return costClient{sqlClient}
}

func (costClient *costClient) GetCost() float64 {
	return 0
}

type Handler struct {
	token      string
	costClient CostClient
}

func Init(account string) Handler {
	return Handler{
		token: account,
	}
}

func (rh Handler) getToken() string {
	return rh.token
}

func (rh Handler) GetCostReport() error {
	costHandler := cost_handler.Init(rh.getToken())
	rh.costClient.ge
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	firstOfLastMonth := firstOfMonth.AddDate(0, -1, +1)
	lastOfLastMonth := firstOfLastMonth.AddDate(0, 1, -1)
	appName := "ole-test"
	cost, err := costHandler.GetTotalCost(&firstOfLastMonth, &lastOfLastMonth, &appName)
}
