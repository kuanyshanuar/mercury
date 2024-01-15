// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/user_cottage.go

// Package mock_domain is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/zharzhanov/mercury/internal/domain"
)

// MockUserCottageRepository is a mock of UserCottageRepository interface.
type MockUserCottageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserCottageRepositoryMockRecorder
}

// MockUserCottageRepositoryMockRecorder is the mock recorder for MockUserCottageRepository.
type MockUserCottageRepositoryMockRecorder struct {
	mock *MockUserCottageRepository
}

// NewMockUserCottageRepository creates a new mock instance.
func NewMockUserCottageRepository(ctrl *gomock.Controller) *MockUserCottageRepository {
	mock := &MockUserCottageRepository{ctrl: ctrl}
	mock.recorder = &MockUserCottageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCottageRepository) EXPECT() *MockUserCottageRepositoryMockRecorder {
	return m.recorder
}

// AddFavouriteCottage mocks base method.
func (m *MockUserCottageRepository) AddFavouriteCottage(ctx context.Context, userID, cottageID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavouriteCottage", ctx, userID, cottageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavouriteCottage indicates an expected call of AddFavouriteCottage.
func (mr *MockUserCottageRepositoryMockRecorder) AddFavouriteCottage(ctx, userID, cottageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavouriteCottage", reflect.TypeOf((*MockUserCottageRepository)(nil).AddFavouriteCottage), ctx, userID, cottageID)
}

// DeleteFavouriteCottage mocks base method.
func (m *MockUserCottageRepository) DeleteFavouriteCottage(ctx context.Context, userID, cottageID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavouriteCottage", ctx, userID, cottageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavouriteCottage indicates an expected call of DeleteFavouriteCottage.
func (mr *MockUserCottageRepositoryMockRecorder) DeleteFavouriteCottage(ctx, userID, cottageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavouriteCottage", reflect.TypeOf((*MockUserCottageRepository)(nil).DeleteFavouriteCottage), ctx, userID, cottageID)
}

// IsCottageExists mocks base method.
func (m *MockUserCottageRepository) IsCottageExists(ctx context.Context, userID, cottageID int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCottageExists", ctx, userID, cottageID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCottageExists indicates an expected call of IsCottageExists.
func (mr *MockUserCottageRepositoryMockRecorder) IsCottageExists(ctx, userID, cottageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCottageExists", reflect.TypeOf((*MockUserCottageRepository)(nil).IsCottageExists), ctx, userID, cottageID)
}

// ListFavouriteCottages mocks base method.
func (m *MockUserCottageRepository) ListFavouriteCottages(ctx context.Context, userID int64, criteria domain.FavouriteCottagesSearchCriteria) ([]int64, domain.Total, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFavouriteCottages", ctx, userID, criteria)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(domain.Total)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFavouriteCottages indicates an expected call of ListFavouriteCottages.
func (mr *MockUserCottageRepositoryMockRecorder) ListFavouriteCottages(ctx, userID, criteria interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFavouriteCottages", reflect.TypeOf((*MockUserCottageRepository)(nil).ListFavouriteCottages), ctx, userID, criteria)
}

// MockUserCottageService is a mock of UserCottageService interface.
type MockUserCottageService struct {
	ctrl     *gomock.Controller
	recorder *MockUserCottageServiceMockRecorder
}

// MockUserCottageServiceMockRecorder is the mock recorder for MockUserCottageService.
type MockUserCottageServiceMockRecorder struct {
	mock *MockUserCottageService
}

// NewMockUserCottageService creates a new mock instance.
func NewMockUserCottageService(ctrl *gomock.Controller) *MockUserCottageService {
	mock := &MockUserCottageService{ctrl: ctrl}
	mock.recorder = &MockUserCottageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCottageService) EXPECT() *MockUserCottageServiceMockRecorder {
	return m.recorder
}

// AddCottageToFavourites mocks base method.
func (m *MockUserCottageService) AddCottageToFavourites(ctx context.Context, userID, CottageID int64, callerID domain.CallerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCottageToFavourites", ctx, userID, CottageID, callerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCottageToFavourites indicates an expected call of AddCottageToFavourites.
func (mr *MockUserCottageServiceMockRecorder) AddCottageToFavourites(ctx, userID, CottageID, callerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCottageToFavourites", reflect.TypeOf((*MockUserCottageService)(nil).AddCottageToFavourites), ctx, userID, CottageID, callerID)
}

// DeleteCottageFromFavourites mocks base method.
func (m *MockUserCottageService) DeleteCottageFromFavourites(ctx context.Context, userID, CottageID int64, callerID domain.CallerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCottageFromFavourites", ctx, userID, CottageID, callerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCottageFromFavourites indicates an expected call of DeleteCottageFromFavourites.
func (mr *MockUserCottageServiceMockRecorder) DeleteCottageFromFavourites(ctx, userID, CottageID, callerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCottageFromFavourites", reflect.TypeOf((*MockUserCottageService)(nil).DeleteCottageFromFavourites), ctx, userID, CottageID, callerID)
}

// ListFavouriteCottages mocks base method.
func (m *MockUserCottageService) ListFavouriteCottages(ctx context.Context, userID int64, criteria domain.FavouriteCottagesSearchCriteria, callerID domain.CallerID) ([]int64, domain.Total, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFavouriteCottages", ctx, userID, criteria, callerID)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(domain.Total)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFavouriteCottages indicates an expected call of ListFavouriteCottages.
func (mr *MockUserCottageServiceMockRecorder) ListFavouriteCottages(ctx, userID, criteria, callerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFavouriteCottages", reflect.TypeOf((*MockUserCottageService)(nil).ListFavouriteCottages), ctx, userID, criteria, callerID)
}