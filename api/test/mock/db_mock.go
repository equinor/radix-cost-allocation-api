// Code generated by MockGen. DO NOT EDIT.
// Source: ./models/sql.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	models "github.com/equinor/radix-cost-allocation-api/models"
	gomock "github.com/golang/mock/gomock"
)

// MockCostRepository is a mock of CostRepository interface.
type MockCostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCostRepositoryMockRecorder
}

// MockCostRepositoryMockRecorder is the mock recorder for MockCostRepository.
type MockCostRepositoryMockRecorder struct {
	mock *MockCostRepository
}

// NewMockCostRepository creates a new mock instance.
func NewMockCostRepository(ctrl *gomock.Controller) *MockCostRepository {
	mock := &MockCostRepository{ctrl: ctrl}
	mock.recorder = &MockCostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCostRepository) EXPECT() *MockCostRepositoryMockRecorder {
	return m.recorder
}

// GetLatestRun mocks base method.
func (m *MockCostRepository) GetLatestRun() (models.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRun")
	ret0, _ := ret[0].(models.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRun indicates an expected call of GetLatestRun.
func (mr *MockCostRepositoryMockRecorder) GetLatestRun() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRun", reflect.TypeOf((*MockCostRepository)(nil).GetLatestRun))
}

// GetRunsBetweenTimes mocks base method.
func (m *MockCostRepository) GetRunsBetweenTimes(from, to *time.Time) ([]models.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunsBetweenTimes", from, to)
	ret0, _ := ret[0].([]models.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunsBetweenTimes indicates an expected call of GetRunsBetweenTimes.
func (mr *MockCostRepositoryMockRecorder) GetRunsBetweenTimes(from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunsBetweenTimes", reflect.TypeOf((*MockCostRepository)(nil).GetRunsBetweenTimes), from, to)
}
