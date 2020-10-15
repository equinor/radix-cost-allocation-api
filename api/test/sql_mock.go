package test

import (
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
)

// FakeDatabase implements Repository interface
type FakeDatabase struct {
}

// GetLatestRun returns static run for testing purposes
func (fdb FakeDatabase) GetLatestRun() (costModels.Run, error) {
	return costModels.Run{
		ClusterMemoryMegaByte: 1000,
		ClusterCPUMillicore:   1000,
	}, nil
}

// GetRunsBetweenTimes returns static runs for testing purposes
func (fdb FakeDatabase) GetRunsBetweenTimes(from, to *time.Time) ([]costModels.Run, error) {
	return nil, nil
}

// CloseDB empty implementation
func (fdb FakeDatabase) CloseDB() {
	return
}

func newFakeDbCon() *FakeDatabase {
	return &FakeDatabase{}
}

// NewFakeCostRepository creates mock CostRepository
func NewFakeCostRepository() *models.CostRepository {
	fakeDbCon := newFakeDbCon()
	return &models.CostRepository{Repo: fakeDbCon}
}
