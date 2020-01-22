// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/api (interfaces: ClientManager)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/vmware-tanzu/octant/internal/api"
	config "github.com/vmware-tanzu/octant/internal/config"
	http "net/http"
	reflect "reflect"
)

// MockClientManager is a mock of ClientManager interface
type MockClientManager struct {
	ctrl     *gomock.Controller
	recorder *MockClientManagerMockRecorder
}

// MockClientManagerMockRecorder is the mock recorder for MockClientManager
type MockClientManagerMockRecorder struct {
	mock *MockClientManager
}

// NewMockClientManager creates a new mock instance
func NewMockClientManager(ctrl *gomock.Controller) *MockClientManager {
	mock := &MockClientManager{ctrl: ctrl}
	mock.recorder = &MockClientManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientManager) EXPECT() *MockClientManagerMockRecorder {
	return m.recorder
}

// ClientFromRequest mocks base method
func (m *MockClientManager) ClientFromRequest(arg0 config.Dash, arg1 http.ResponseWriter, arg2 *http.Request) (*api.WebsocketClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientFromRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.WebsocketClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientFromRequest indicates an expected call of ClientFromRequest
func (mr *MockClientManagerMockRecorder) ClientFromRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientFromRequest", reflect.TypeOf((*MockClientManager)(nil).ClientFromRequest), arg0, arg1, arg2)
}

// Clients mocks base method
func (m *MockClientManager) Clients() []*api.WebsocketClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clients")
	ret0, _ := ret[0].([]*api.WebsocketClient)
	return ret0
}

// Clients indicates an expected call of Clients
func (mr *MockClientManagerMockRecorder) Clients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clients", reflect.TypeOf((*MockClientManager)(nil).Clients))
}

// Run mocks base method
func (m *MockClientManager) Run(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", arg0)
}

// Run indicates an expected call of Run
func (mr *MockClientManagerMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockClientManager)(nil).Run), arg0)
}
