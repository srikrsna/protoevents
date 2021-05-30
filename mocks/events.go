// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/srikrsna/protoevents (interfaces: Dispatcher,ErrorLogger)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protoevents "github.com/srikrsna/protoevents"
)

// MockDispatcher is a mock of Dispatcher interface.
type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

// MockDispatcherMockRecorder is the mock recorder for MockDispatcher.
type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

// NewMockDispatcher creates a new mock instance.
func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

// Dispatch mocks base method.
func (m *MockDispatcher) Dispatch(arg0 context.Context, arg1 *protoevents.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispatch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockDispatcherMockRecorder) Dispatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockDispatcher)(nil).Dispatch), arg0, arg1)
}

// MockErrorLogger is a mock of ErrorLogger interface.
type MockErrorLogger struct {
	ctrl     *gomock.Controller
	recorder *MockErrorLoggerMockRecorder
}

// MockErrorLoggerMockRecorder is the mock recorder for MockErrorLogger.
type MockErrorLoggerMockRecorder struct {
	mock *MockErrorLogger
}

// NewMockErrorLogger creates a new mock instance.
func NewMockErrorLogger(ctrl *gomock.Controller) *MockErrorLogger {
	mock := &MockErrorLogger{ctrl: ctrl}
	mock.recorder = &MockErrorLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockErrorLogger) EXPECT() *MockErrorLoggerMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockErrorLogger) Log(arg0 context.Context, arg1 *protoevents.Event, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", arg0, arg1, arg2)
}

// Log indicates an expected call of Log.
func (mr *MockErrorLoggerMockRecorder) Log(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockErrorLogger)(nil).Log), arg0, arg1, arg2)
}
