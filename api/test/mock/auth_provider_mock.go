// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/utils/auth/auth_provider.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	auth "github.com/equinor/radix-cost-allocation-api/api/utils/auth"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthProvider is a mock of AuthProvider interface.
type MockAuthProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAuthProviderMockRecorder
}

// MockAuthProviderMockRecorder is the mock recorder for MockAuthProvider.
type MockAuthProviderMockRecorder struct {
	mock *MockAuthProvider
}

// NewMockAuthProvider creates a new mock instance.
func NewMockAuthProvider(ctrl *gomock.Controller) *MockAuthProvider {
	mock := &MockAuthProvider{ctrl: ctrl}
	mock.recorder = &MockAuthProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthProvider) EXPECT() *MockAuthProviderMockRecorder {
	return m.recorder
}

// VerifyToken mocks base method.
func (m *MockAuthProvider) VerifyToken(ctx context.Context, token string) (auth.IDToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", ctx, token)
	ret0, _ := ret[0].(auth.IDToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockAuthProviderMockRecorder) VerifyToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockAuthProvider)(nil).VerifyToken), ctx, token)
}

// MockIDToken is a mock of IDToken interface.
type MockIDToken struct {
	ctrl     *gomock.Controller
	recorder *MockIDTokenMockRecorder
}

// MockIDTokenMockRecorder is the mock recorder for MockIDToken.
type MockIDTokenMockRecorder struct {
	mock *MockIDToken
}

// NewMockIDToken creates a new mock instance.
func NewMockIDToken(ctrl *gomock.Controller) *MockIDToken {
	mock := &MockIDToken{ctrl: ctrl}
	mock.recorder = &MockIDTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDToken) EXPECT() *MockIDTokenMockRecorder {
	return m.recorder
}

// GetClaims mocks base method.
func (m *MockIDToken) GetClaims(out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaims", out)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetClaims indicates an expected call of GetClaims.
func (mr *MockIDTokenMockRecorder) GetClaims(out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaims", reflect.TypeOf((*MockIDToken)(nil).GetClaims), out)
}
