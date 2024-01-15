// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/zharzhanov/mercury/internal/domain (interfaces: CodeRedisRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/zharzhanov/mercury/internal/domain"
)

// MockCodeRedisRepository is a mock of CodeRedisRepository interface.
type MockCodeRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCodeRedisRepositoryMockRecorder
}

// MockCodeRedisRepositoryMockRecorder is the mock recorder for MockCodeRedisRepository.
type MockCodeRedisRepositoryMockRecorder struct {
	mock *MockCodeRedisRepository
}

// NewMockCodeRedisRepository creates a new mock instance.
func NewMockCodeRedisRepository(ctrl *gomock.Controller) *MockCodeRedisRepository {
	mock := &MockCodeRedisRepository{ctrl: ctrl}
	mock.recorder = &MockCodeRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeRedisRepository) EXPECT() *MockCodeRedisRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCodeRedisRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCodeRedisRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCodeRedisRepository)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockCodeRedisRepository) Get(arg0 context.Context, arg1 string) (*domain.Code, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*domain.Code)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCodeRedisRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCodeRedisRepository)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockCodeRedisRepository) Set(arg0 context.Context, arg1 string, arg2 *domain.Code, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCodeRedisRepositoryMockRecorder) Set(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCodeRedisRepository)(nil).Set), arg0, arg1, arg2, arg3)
}
