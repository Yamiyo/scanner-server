// Code generated by MockGen. DO NOT EDIT.
// Source: blocks_repo.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	po "portto-homework/internal/model/po"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockBlocksRepo is a mock of BlocksRepo interface.
type MockBlocksRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBlocksRepoMockRecorder
}

// MockBlocksRepoMockRecorder is the mock recorder for MockBlocksRepo.
type MockBlocksRepoMockRecorder struct {
	mock *MockBlocksRepo
}

// NewMockBlocksRepo creates a new mock instance.
func NewMockBlocksRepo(ctrl *gomock.Controller) *MockBlocksRepo {
	mock := &MockBlocksRepo{ctrl: ctrl}
	mock.recorder = &MockBlocksRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlocksRepo) EXPECT() *MockBlocksRepoMockRecorder {
	return m.recorder
}

// GetBlockByNum mocks base method.
func (m *MockBlocksRepo) GetBlockByNum(ctx context.Context, db *gorm.DB, num uint64) (*po.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByNum", ctx, db, num)
	ret0, _ := ret[0].(*po.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNum indicates an expected call of GetBlockByNum.
func (mr *MockBlocksRepoMockRecorder) GetBlockByNum(ctx, db, num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNum", reflect.TypeOf((*MockBlocksRepo)(nil).GetBlockByNum), ctx, db, num)
}

// GetBlockLatestN mocks base method.
func (m *MockBlocksRepo) GetBlockLatestN(ctx context.Context, db *gorm.DB, n int) ([]*po.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockLatestN", ctx, db, n)
	ret0, _ := ret[0].([]*po.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockLatestN indicates an expected call of GetBlockLatestN.
func (mr *MockBlocksRepoMockRecorder) GetBlockLatestN(ctx, db, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockLatestN", reflect.TypeOf((*MockBlocksRepo)(nil).GetBlockLatestN), ctx, db, n)
}
