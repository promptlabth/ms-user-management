// Code generated by MockGen. DO NOT EDIT.
// Source: port.go
//
// Generated by this command:
//
//	mockgen -source=port.go -package=user -destination=mock_port_test.go
//

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserServicer is a mock of UserServicer interface.
type MockUserServicer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServicerMockRecorder
}

// MockUserServicerMockRecorder is the mock recorder for MockUserServicer.
type MockUserServicerMockRecorder struct {
	mock *MockUserServicer
}

// NewMockUserServicer creates a new mock instance.
func NewMockUserServicer(ctrl *gomock.Controller) *MockUserServicer {
	mock := &MockUserServicer{ctrl: ctrl}
	mock.recorder = &MockUserServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServicer) EXPECT() *MockUserServicerMockRecorder {
	return m.recorder
}

// GetUserById mocks base method.
func (m *MockUserServicer) GetUserById() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUserById")
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServicerMockRecorder) GetUserById() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserServicer)(nil).GetUserById))
}

// MockuserRepository is a mock of userRepository interface.
type MockuserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockuserRepositoryMockRecorder
}

// MockuserRepositoryMockRecorder is the mock recorder for MockuserRepository.
type MockuserRepositoryMockRecorder struct {
	mock *MockuserRepository
}

// NewMockuserRepository creates a new mock instance.
func NewMockuserRepository(ctrl *gomock.Controller) *MockuserRepository {
	mock := &MockuserRepository{ctrl: ctrl}
	mock.recorder = &MockuserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserRepository) EXPECT() *MockuserRepositoryMockRecorder {
	return m.recorder
}

// CreateUserBalance mocks base method.
func (m *MockuserRepository) CreateUserBalance(ctx context.Context, tx *gorm.DB, userBalanceEntity *UserBalanceMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserBalance", ctx, tx, userBalanceEntity)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserBalance indicates an expected call of CreateUserBalance.
func (mr *MockuserRepositoryMockRecorder) CreateUserBalance(ctx, tx, userBalanceEntity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserBalance", reflect.TypeOf((*MockuserRepository)(nil).CreateUserBalance), ctx, tx, userBalanceEntity)
}

// GetDetailUserById mocks base method.
func (m *MockuserRepository) GetDetailUserById(ctx context.Context, id string) (*UserDetailInquirey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailUserById", ctx, id)
	ret0, _ := ret[0].(*UserDetailInquirey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailUserById indicates an expected call of GetDetailUserById.
func (mr *MockuserRepositoryMockRecorder) GetDetailUserById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailUserById", reflect.TypeOf((*MockuserRepository)(nil).GetDetailUserById), ctx, id)
}

// GetPlanById mocks base method.
func (m *MockuserRepository) GetPlanById(ctx context.Context, planId int64) (*PlanEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanById", ctx, planId)
	ret0, _ := ret[0].(*PlanEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlanById indicates an expected call of GetPlanById.
func (mr *MockuserRepositoryMockRecorder) GetPlanById(ctx, planId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanById", reflect.TypeOf((*MockuserRepository)(nil).GetPlanById), ctx, planId)
}

// GetPlanByType mocks base method.
func (m *MockuserRepository) GetPlanByType(ctx context.Context, planType string) (*PlanEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanByType", ctx, planType)
	ret0, _ := ret[0].(*PlanEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlanByType indicates an expected call of GetPlanByType.
func (mr *MockuserRepositoryMockRecorder) GetPlanByType(ctx, planType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanByType", reflect.TypeOf((*MockuserRepository)(nil).GetPlanByType), ctx, planType)
}

// Transactional mocks base method.
func (m *MockuserRepository) Transactional(fn func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transactional", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transactional indicates an expected call of Transactional.
func (mr *MockuserRepositoryMockRecorder) Transactional(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transactional", reflect.TypeOf((*MockuserRepository)(nil).Transactional), fn)
}

// UpsertUser mocks base method.
func (m *MockuserRepository) UpsertUser(ctx context.Context, tx *gorm.DB, userEntity *UserEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertUser", ctx, tx, userEntity)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertUser indicates an expected call of UpsertUser.
func (mr *MockuserRepositoryMockRecorder) UpsertUser(ctx, tx, userEntity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertUser", reflect.TypeOf((*MockuserRepository)(nil).UpsertUser), ctx, tx, userEntity)
}
