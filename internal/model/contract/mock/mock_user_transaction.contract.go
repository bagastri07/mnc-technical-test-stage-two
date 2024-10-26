// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/model/contract/user_transaction.contract.go

// Package mock_contract is a generated GoMock package.
package mock_contract

import (
	context "context"
	reflect "reflect"

	entity "github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserTransactionRepository is a mock of UserTransactionRepository interface.
type MockUserTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserTransactionRepositoryMockRecorder
}

// MockUserTransactionRepositoryMockRecorder is the mock recorder for MockUserTransactionRepository.
type MockUserTransactionRepositoryMockRecorder struct {
	mock *MockUserTransactionRepository
}

// NewMockUserTransactionRepository creates a new mock instance.
func NewMockUserTransactionRepository(ctrl *gomock.Controller) *MockUserTransactionRepository {
	mock := &MockUserTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockUserTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserTransactionRepository) EXPECT() *MockUserTransactionRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockUserTransactionRepository) FindByID(ctx context.Context, ID uuid.UUID) (*entity.UserTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, ID)
	ret0, _ := ret[0].(*entity.UserTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserTransactionRepositoryMockRecorder) FindByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserTransactionRepository)(nil).FindByID), ctx, ID)
}

// GetByUserID mocks base method.
func (m *MockUserTransactionRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.UserTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", ctx, userID)
	ret0, _ := ret[0].([]entity.UserTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID.
func (mr *MockUserTransactionRepositoryMockRecorder) GetByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockUserTransactionRepository)(nil).GetByUserID), ctx, userID)
}

// Upsert mocks base method.
func (m *MockUserTransactionRepository) Upsert(ctx context.Context, userTransaction *entity.UserTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, userTransaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockUserTransactionRepositoryMockRecorder) Upsert(ctx, userTransaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockUserTransactionRepository)(nil).Upsert), ctx, userTransaction)
}
