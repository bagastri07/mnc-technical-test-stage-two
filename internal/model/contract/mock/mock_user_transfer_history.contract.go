// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/model/contract/user_transfer_history.contract.go

// Package mock_contract is a generated GoMock package.
package mock_contract

import (
	context "context"
	reflect "reflect"

	entity "github.com/bagastri07/mnc-technical-test-stage-two/internal/model/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockUserTransferHistoryRepository is a mock of UserTransferHistoryRepository interface.
type MockUserTransferHistoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserTransferHistoryRepositoryMockRecorder
}

// MockUserTransferHistoryRepositoryMockRecorder is the mock recorder for MockUserTransferHistoryRepository.
type MockUserTransferHistoryRepositoryMockRecorder struct {
	mock *MockUserTransferHistoryRepository
}

// NewMockUserTransferHistoryRepository creates a new mock instance.
func NewMockUserTransferHistoryRepository(ctrl *gomock.Controller) *MockUserTransferHistoryRepository {
	mock := &MockUserTransferHistoryRepository{ctrl: ctrl}
	mock.recorder = &MockUserTransferHistoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserTransferHistoryRepository) EXPECT() *MockUserTransferHistoryRepositoryMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockUserTransferHistoryRepository) Upsert(ctx context.Context, userTransferHistory *entity.UserTransferHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, userTransferHistory)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockUserTransferHistoryRepositoryMockRecorder) Upsert(ctx, userTransferHistory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockUserTransferHistoryRepository)(nil).Upsert), ctx, userTransferHistory)
}
