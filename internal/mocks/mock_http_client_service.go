// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/zharzhanov/mercury/internal/domain (interfaces: HTTPClientService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHTTPClientService is a mock of HTTPClientService interface.
type MockHTTPClientService struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientServiceMockRecorder
}

// MockHTTPClientServiceMockRecorder is the mock recorder for MockHTTPClientService.
type MockHTTPClientServiceMockRecorder struct {
	mock *MockHTTPClientService
}

// NewMockHTTPClientService creates a new mock instance.
func NewMockHTTPClientService(ctrl *gomock.Controller) *MockHTTPClientService {
	mock := &MockHTTPClientService{ctrl: ctrl}
	mock.recorder = &MockHTTPClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClientService) EXPECT() *MockHTTPClientServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHTTPClientService) Get(arg0 context.Context, arg1 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHTTPClientServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPClientService)(nil).Get), arg0, arg1)
}

// Post mocks base method.
func (m *MockHTTPClientService) Post(arg0 context.Context, arg1 string, arg2 interface{}) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockHTTPClientServiceMockRecorder) Post(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHTTPClientService)(nil).Post), arg0, arg1, arg2)
}
