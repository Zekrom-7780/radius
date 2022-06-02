// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/queue (interfaces: Dequeuer)

// Package queue is a generated GoMock package.
package queue

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDequeuer is a mock of Dequeuer interface.
type MockDequeuer struct {
	ctrl     *gomock.Controller
	recorder *MockDequeuerMockRecorder
}

// MockDequeuerMockRecorder is the mock recorder for MockDequeuer.
type MockDequeuerMockRecorder struct {
	mock *MockDequeuer
}

// NewMockDequeuer creates a new mock instance.
func NewMockDequeuer(ctrl *gomock.Controller) *MockDequeuer {
	mock := &MockDequeuer{ctrl: ctrl}
	mock.recorder = &MockDequeuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDequeuer) EXPECT() *MockDequeuerMockRecorder {
	return m.recorder
}

// Dequeue mocks base method.
func (m *MockDequeuer) Dequeue(arg0 context.Context, arg1 ...DequeueOptions) (<-chan *Message, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Dequeue", varargs...)
	ret0, _ := ret[0].(<-chan *Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dequeue indicates an expected call of Dequeue.
func (mr *MockDequeuerMockRecorder) Dequeue(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dequeue", reflect.TypeOf((*MockDequeuer)(nil).Dequeue), varargs...)
}