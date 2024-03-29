// Code generated by MockGen. DO NOT EDIT.
// Source: transaction_repo.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	po "portto-homework/internal/model/po"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockTransactionsRepo is a mock of TransactionsRepo interface.
type MockTransactionsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionsRepoMockRecorder
}

// MockTransactionsRepoMockRecorder is the mock recorder for MockTransactionsRepo.
type MockTransactionsRepoMockRecorder struct {
	mock *MockTransactionsRepo
}

// NewMockTransactionsRepo creates a new mock instance.
func NewMockTransactionsRepo(ctrl *gomock.Controller) *MockTransactionsRepo {
	mock := &MockTransactionsRepo{ctrl: ctrl}
	mock.recorder = &MockTransactionsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionsRepo) EXPECT() *MockTransactionsRepoMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionsRepo) CreateTransaction(ctx context.Context, db *gorm.DB, data []*po.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, db, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionsRepoMockRecorder) CreateTransaction(ctx, db, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionsRepo)(nil).CreateTransaction), ctx, db, data)
}

// CreateTransactionLogs mocks base method.
func (m *MockTransactionsRepo) CreateTransactionLogs(ctx context.Context, db *gorm.DB, data []*po.TransactionLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransactionLogs", ctx, db, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTransactionLogs indicates an expected call of CreateTransactionLogs.
func (mr *MockTransactionsRepoMockRecorder) CreateTransactionLogs(ctx, db, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransactionLogs", reflect.TypeOf((*MockTransactionsRepo)(nil).CreateTransactionLogs), ctx, db, data)
}
