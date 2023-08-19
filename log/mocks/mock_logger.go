// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/casbin/casbin/v2/log (interfaces: Logger)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// EnableLog mocks base method.
func (m *MockLogger) EnableLog(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnableLog", arg0)
}

// EnableLog indicates an expected call of EnableLog.
func (mr *MockLoggerMockRecorder) EnableLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableLog", reflect.TypeOf((*MockLogger)(nil).EnableLog), arg0)
}

// IsEnabled mocks base method.
func (m *MockLogger) IsEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEnabled indicates an expected call of IsEnabled.
func (mr *MockLoggerMockRecorder) IsEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabled", reflect.TypeOf((*MockLogger)(nil).IsEnabled))
}

// LogEnforce mocks base method.
func (m *MockLogger) LogEnforce(arg0 string, arg1 []interface{}, arg2 bool, arg3 [][]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogEnforce", arg0, arg1, arg2, arg3)
}

// LogEnforce indicates an expected call of LogEnforce.
func (mr *MockLoggerMockRecorder) LogEnforce(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogEnforce", reflect.TypeOf((*MockLogger)(nil).LogEnforce), arg0, arg1, arg2, arg3)
}

// LogError mocks base method.
func (m *MockLogger) LogError(arg0 error, arg1 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogError", varargs...)
}

// LogError indicates an expected call of LogError.
func (mr *MockLoggerMockRecorder) LogError(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockLogger)(nil).LogError), varargs...)
}

// LogModel mocks base method.
func (m *MockLogger) LogModel(arg0 [][]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogModel", arg0)
}

// LogModel indicates an expected call of LogModel.
func (mr *MockLoggerMockRecorder) LogModel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogModel", reflect.TypeOf((*MockLogger)(nil).LogModel), arg0)
}

// LogPolicy mocks base method.
func (m *MockLogger) LogPolicy(arg0 map[string][][]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogPolicy", arg0)
}

// LogPolicy indicates an expected call of LogPolicy.
func (mr *MockLoggerMockRecorder) LogPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogPolicy", reflect.TypeOf((*MockLogger)(nil).LogPolicy), arg0)
}

// LogRole mocks base method.
func (m *MockLogger) LogRole(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogRole", arg0)
}

// LogRole indicates an expected call of LogRole.
func (mr *MockLoggerMockRecorder) LogRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogRole", reflect.TypeOf((*MockLogger)(nil).LogRole), arg0)
}
