// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/zharzhanov/mercury/internal/domain (interfaces: LeadBuilderService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/zharzhanov/mercury/internal/domain"
)

// MockLeadBuilderService is a mock of LeadBuilderService interface.
type MockLeadBuilderService struct {
	ctrl     *gomock.Controller
	recorder *MockLeadBuilderServiceMockRecorder
}

// MockLeadBuilderServiceMockRecorder is the mock recorder for MockLeadBuilderService.
type MockLeadBuilderServiceMockRecorder struct {
	mock *MockLeadBuilderService
}

// NewMockLeadBuilderService creates a new mock instance.
func NewMockLeadBuilderService(ctrl *gomock.Controller) *MockLeadBuilderService {
	mock := &MockLeadBuilderService{ctrl: ctrl}
	mock.recorder = &MockLeadBuilderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeadBuilderService) EXPECT() *MockLeadBuilderServiceMockRecorder {
	return m.recorder
}

// CreateLeadBuilder mocks base method.
func (m *MockLeadBuilderService) CreateLeadBuilder(arg0 context.Context, arg1 *domain.LeadBuilder, arg2 domain.CallerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLeadBuilder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLeadBuilder indicates an expected call of CreateLeadBuilder.
func (mr *MockLeadBuilderServiceMockRecorder) CreateLeadBuilder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLeadBuilder", reflect.TypeOf((*MockLeadBuilderService)(nil).CreateLeadBuilder), arg0, arg1, arg2)
}

// DeleteLeadBuilder mocks base method.
func (m *MockLeadBuilderService) DeleteLeadBuilder(arg0 context.Context, arg1 domain.LeadID, arg2 domain.CallerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLeadBuilder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLeadBuilder indicates an expected call of DeleteLeadBuilder.
func (mr *MockLeadBuilderServiceMockRecorder) DeleteLeadBuilder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLeadBuilder", reflect.TypeOf((*MockLeadBuilderService)(nil).DeleteLeadBuilder), arg0, arg1, arg2)
}

// GetLeadBuilder mocks base method.
func (m *MockLeadBuilderService) GetLeadBuilder(arg0 context.Context, arg1 domain.LeadID, arg2 domain.CallerID) (*domain.LeadBuilder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeadBuilder", arg0, arg1, arg2)
	ret0, _ := ret[0].(*domain.LeadBuilder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeadBuilder indicates an expected call of GetLeadBuilder.
func (mr *MockLeadBuilderServiceMockRecorder) GetLeadBuilder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeadBuilder", reflect.TypeOf((*MockLeadBuilderService)(nil).GetLeadBuilder), arg0, arg1, arg2)
}

// ListLeadBuilders mocks base method.
func (m *MockLeadBuilderService) ListLeadBuilders(arg0 context.Context, arg1 domain.LeadBuilderSearchCriteria, arg2 domain.CallerID) ([]*domain.LeadBuilder, domain.Total, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLeadBuilders", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*domain.LeadBuilder)
	ret1, _ := ret[1].(domain.Total)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListLeadBuilders indicates an expected call of ListLeadBuilders.
func (mr *MockLeadBuilderServiceMockRecorder) ListLeadBuilders(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLeadBuilders", reflect.TypeOf((*MockLeadBuilderService)(nil).ListLeadBuilders), arg0, arg1, arg2)
}

// RevokeLeadBuilder mocks base method.
func (m *MockLeadBuilderService) RevokeLeadBuilder(arg0 context.Context, arg1 domain.LeadID, arg2 domain.CallerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeLeadBuilder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeLeadBuilder indicates an expected call of RevokeLeadBuilder.
func (mr *MockLeadBuilderServiceMockRecorder) RevokeLeadBuilder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeLeadBuilder", reflect.TypeOf((*MockLeadBuilderService)(nil).RevokeLeadBuilder), arg0, arg1, arg2)
}

// UpdateLeadBuilder mocks base method.
func (m *MockLeadBuilderService) UpdateLeadBuilder(arg0 context.Context, arg1 domain.LeadID, arg2 *domain.LeadBuilder, arg3 domain.CallerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLeadBuilder", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLeadBuilder indicates an expected call of UpdateLeadBuilder.
func (mr *MockLeadBuilderServiceMockRecorder) UpdateLeadBuilder(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLeadBuilder", reflect.TypeOf((*MockLeadBuilderService)(nil).UpdateLeadBuilder), arg0, arg1, arg2, arg3)
}
