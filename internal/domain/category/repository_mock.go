// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/category/repository.go

// Package category is a generated GoMock package.
package category

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// BindOrganizations mocks base method.
func (m *MockRepository) BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindOrganizations", ctx, buildingID, organizationIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindOrganizations indicates an expected call of BindOrganizations.
func (mr *MockRepositoryMockRecorder) BindOrganizations(ctx, buildingID, organizationIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindOrganizations", reflect.TypeOf((*MockRepository)(nil).BindOrganizations), ctx, buildingID, organizationIDs)
}

// Count mocks base method.
func (m *MockRepository) Count(ctx context.Context, cond *QueryConditions) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, cond)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRepositoryMockRecorder) Count(ctx, cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRepository)(nil).Count), ctx, cond)
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, c *Category) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, c)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, c)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, id)
}

// First mocks base method.
func (m *MockRepository) First(ctx context.Context, cond *Category) (*Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First", ctx, cond)
	ret0, _ := ret[0].(*Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// First indicates an expected call of First.
func (mr *MockRepositoryMockRecorder) First(ctx, cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockRepository)(nil).First), ctx, cond)
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, id uint) (*Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, id)
}

// Query mocks base method.
func (m *MockRepository) Query(ctx context.Context, cond *QueryConditions) ([]Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, cond)
	ret0, _ := ret[0].([]Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockRepositoryMockRecorder) Query(ctx, cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRepository)(nil).Query), ctx, cond)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, c *Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, c)
}