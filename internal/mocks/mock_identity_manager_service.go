// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/zharzhanov/mercury/internal/domain (interfaces: IdentityManagerService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/zharzhanov/mercury/internal/domain"
)

// MockIdentityManagerService is a mock of IdentityManagerService interface.
type MockIdentityManagerService struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityManagerServiceMockRecorder
}

// MockIdentityManagerServiceMockRecorder is the mock recorder for MockIdentityManagerService.
type MockIdentityManagerServiceMockRecorder struct {
	mock *MockIdentityManagerService
}

// NewMockIdentityManagerService creates a new mock instance.
func NewMockIdentityManagerService(ctrl *gomock.Controller) *MockIdentityManagerService {
	mock := &MockIdentityManagerService{ctrl: ctrl}
	mock.recorder = &MockIdentityManagerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityManagerService) EXPECT() *MockIdentityManagerServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIdentityManagerService) CreateUser(arg0 context.Context, arg1 *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIdentityManagerServiceMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIdentityManagerService)(nil).CreateUser), arg0, arg1)
}

// ResetPassword mocks base method.
func (m *MockIdentityManagerService) ResetPassword(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockIdentityManagerServiceMockRecorder) ResetPassword(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockIdentityManagerService)(nil).ResetPassword), arg0, arg1, arg2)
}

// SendResetPasswordToken mocks base method.
func (m *MockIdentityManagerService) SendResetPasswordToken(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendResetPasswordToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendResetPasswordToken indicates an expected call of SendResetPasswordToken.
func (mr *MockIdentityManagerServiceMockRecorder) SendResetPasswordToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResetPasswordToken", reflect.TypeOf((*MockIdentityManagerService)(nil).SendResetPasswordToken), arg0, arg1)
}

// ValidateCode mocks base method.
func (m *MockIdentityManagerService) ValidateCode(arg0 context.Context, arg1 string) (domain.UserID, domain.RoleID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCode", arg0, arg1)
	ret0, _ := ret[0].(domain.UserID)
	ret1, _ := ret[1].(domain.RoleID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateCode indicates an expected call of ValidateCode.
func (mr *MockIdentityManagerServiceMockRecorder) ValidateCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCode", reflect.TypeOf((*MockIdentityManagerService)(nil).ValidateCode), arg0, arg1)
}

// ValidateUser mocks base method.
func (m *MockIdentityManagerService) ValidateUser(arg0 context.Context, arg1, arg2 string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUser indicates an expected call of ValidateUser.
func (mr *MockIdentityManagerServiceMockRecorder) ValidateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUser", reflect.TypeOf((*MockIdentityManagerService)(nil).ValidateUser), arg0, arg1, arg2)
}