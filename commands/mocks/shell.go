// Code generated by MockGen. DO NOT EDIT.
// Source: shell.go

// Package mock_commands is a generated GoMock package.
package mock_commands

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockShellCommander is a mock of ShellCommander interface.
type MockShellCommander struct {
	ctrl     *gomock.Controller
	recorder *MockShellCommanderMockRecorder
}

// MockShellCommanderMockRecorder is the mock recorder for MockShellCommander.
type MockShellCommanderMockRecorder struct {
	mock *MockShellCommander
}

// NewMockShellCommander creates a new mock instance.
func NewMockShellCommander(ctrl *gomock.Controller) *MockShellCommander {
	mock := &MockShellCommander{ctrl: ctrl}
	mock.recorder = &MockShellCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShellCommander) EXPECT() *MockShellCommanderMockRecorder {
	return m.recorder
}

// RunCommand mocks base method.
func (m *MockShellCommander) RunCommand(name string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommand", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCommand indicates an expected call of RunCommand.
func (mr *MockShellCommanderMockRecorder) RunCommand(name interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockShellCommander)(nil).RunCommand), varargs...)
}
