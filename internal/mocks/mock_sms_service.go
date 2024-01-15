// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/zharzhanov/mercury/internal/domain (interfaces: SmsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/zharzhanov/mercury/internal/domain"
)

// MockSmsService is a mock of SmsService interface.
type MockSmsService struct {
	ctrl     *gomock.Controller
	recorder *MockSmsServiceMockRecorder
}

// MockSmsServiceMockRecorder is the mock recorder for MockSmsService.
type MockSmsServiceMockRecorder struct {
	mock *MockSmsService
}

// NewMockSmsService creates a new mock instance.
func NewMockSmsService(ctrl *gomock.Controller) *MockSmsService {
	mock := &MockSmsService{ctrl: ctrl}
	mock.recorder = &MockSmsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSmsService) EXPECT() *MockSmsServiceMockRecorder {
	return m.recorder
}

// SendSms mocks base method.
func (m *MockSmsService) SendSms(arg0 context.Context, arg1 *domain.Sms) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSms", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSms indicates an expected call of SendSms.
func (mr *MockSmsServiceMockRecorder) SendSms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSms", reflect.TypeOf((*MockSmsService)(nil).SendSms), arg0, arg1)
}
