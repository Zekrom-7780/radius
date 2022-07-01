// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/corerp/backend/deployment (interfaces: DeploymentProcessor)

// Package deployment is a generated GoMock package.
package deployment

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	conv "github.com/project-radius/radius/pkg/armrpc/api/conv"
	renderers "github.com/project-radius/radius/pkg/corerp/renderers"
	outputresource "github.com/project-radius/radius/pkg/radrp/outputresource"
	resources "github.com/project-radius/radius/pkg/ucp/resources"
)

// MockDeploymentProcessor is a mock of DeploymentProcessor interface.
type MockDeploymentProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentProcessorMockRecorder
}

// MockDeploymentProcessorMockRecorder is the mock recorder for MockDeploymentProcessor.
type MockDeploymentProcessorMockRecorder struct {
	mock *MockDeploymentProcessor
}

// NewMockDeploymentProcessor creates a new mock instance.
func NewMockDeploymentProcessor(ctrl *gomock.Controller) *MockDeploymentProcessor {
	mock := &MockDeploymentProcessor{ctrl: ctrl}
	mock.recorder = &MockDeploymentProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentProcessor) EXPECT() *MockDeploymentProcessorMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeploymentProcessor) Delete(arg0 context.Context, arg1 resources.ID, arg2 []outputresource.OutputResource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeploymentProcessorMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeploymentProcessor)(nil).Delete), arg0, arg1, arg2)
}

// Deploy mocks base method.
func (m *MockDeploymentProcessor) Deploy(arg0 context.Context, arg1 resources.ID, arg2 renderers.RendererOutput) (DeploymentOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0, arg1, arg2)
	ret0, _ := ret[0].(DeploymentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deploy indicates an expected call of Deploy.
func (mr *MockDeploymentProcessorMockRecorder) Deploy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeploymentProcessor)(nil).Deploy), arg0, arg1, arg2)
}

// FetchSecrets mocks base method.
func (m *MockDeploymentProcessor) FetchSecrets(arg0 context.Context, arg1 ResourceData) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSecrets", arg0, arg1)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSecrets indicates an expected call of FetchSecrets.
func (mr *MockDeploymentProcessorMockRecorder) FetchSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSecrets", reflect.TypeOf((*MockDeploymentProcessor)(nil).FetchSecrets), arg0, arg1)
}

// Render mocks base method.
func (m *MockDeploymentProcessor) Render(arg0 context.Context, arg1 resources.ID, arg2 conv.DataModelInterface) (renderers.RendererOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1, arg2)
	ret0, _ := ret[0].(renderers.RendererOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render.
func (mr *MockDeploymentProcessorMockRecorder) Render(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockDeploymentProcessor)(nil).Render), arg0, arg1, arg2)
}
