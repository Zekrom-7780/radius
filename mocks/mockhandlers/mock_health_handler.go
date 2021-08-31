// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/radius/pkg/handlers (interfaces: HealthHandler)

// Package mockhandlers is a generated GoMock package.
package mockhandlers

import (
	context "context"
	reflect "reflect"

	healthcontract "github.com/Azure/radius/pkg/healthcontract"
	gomock "github.com/golang/mock/gomock"
)

// MockHealthHandler is a mock of HealthHandler interface.
type MockHealthHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHealthHandlerMockRecorder
}

// MockHealthHandlerMockRecorder is the mock recorder for MockHealthHandler.
type MockHealthHandlerMockRecorder struct {
	mock *MockHealthHandler
}

// NewMockHealthHandler creates a new mock instance.
func NewMockHealthHandler(ctrl *gomock.Controller) *MockHealthHandler {
	mock := &MockHealthHandler{ctrl: ctrl}
	mock.recorder = &MockHealthHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthHandler) EXPECT() *MockHealthHandlerMockRecorder {
	return m.recorder
}

// GetHealthOptions mocks base method.
func (m *MockHealthHandler) GetHealthOptions(arg0 context.Context) healthcontract.HealthCheckOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHealthOptions", arg0)
	ret0, _ := ret[0].(healthcontract.HealthCheckOptions)
	return ret0
}

// GetHealthOptions indicates an expected call of GetHealthOptions.
func (mr *MockHealthHandlerMockRecorder) GetHealthOptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHealthOptions", reflect.TypeOf((*MockHealthHandler)(nil).GetHealthOptions), arg0)
}