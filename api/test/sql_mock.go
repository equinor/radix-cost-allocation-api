package test

import (
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
	models "github.com/equinor/radix-cost-allocation-api/models"
)

type FakeDatabase struct {
}

func (fdb FakeDatabase) GetLatestRun() (costModels.Run, error) {
	return costModels.Run{
		ClusterMemoryMegaByte: 1000,
		ClusterCPUMillicore:   1000,
	}, nil
}

func (fdb FakeDatabase) GetRunsBetweenTimes(from, to *time.Time) ([]costModels.Run, error) {
	return nil, nil
}

func (fdb FakeDatabase) CloseDB() {
	return
}

func newFakeDbCon() *FakeDatabase {
	return &FakeDatabase{}
}

func NewFakeCostRepository() *models.CostRepository {
	fakeDbCon := newFakeDbCon()
	return &models.CostRepository{fakeDbCon}
}
