// Code generated by MockGen. DO NOT EDIT.
// Source: transaction_ctrl.go

// Package mock_restctl is a generated GoMock package.
package mock_restctl

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockTxnCtrl is a mock of TxnCtrl interface.
type MockTxnCtrl struct {
	ctrl     *gomock.Controller
	recorder *MockTxnCtrlMockRecorder
}

// MockTxnCtrlMockRecorder is the mock recorder for MockTxnCtrl.
type MockTxnCtrlMockRecorder struct {
	mock *MockTxnCtrl
}

// NewMockTxnCtrl creates a new mock instance.
func NewMockTxnCtrl(ctrl *gomock.Controller) *MockTxnCtrl {
	mock := &MockTxnCtrl{ctrl: ctrl}
	mock.recorder = &MockTxnCtrlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxnCtrl) EXPECT() *MockTxnCtrlMockRecorder {
	return m.recorder
}

// GetTxnInfo mocks base method.
func (m *MockTxnCtrl) GetTxnInfo(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetTxnInfo", ctx)
}

// GetTxnInfo indicates an expected call of GetTxnInfo.
func (mr *MockTxnCtrlMockRecorder) GetTxnInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxnInfo", reflect.TypeOf((*MockTxnCtrl)(nil).GetTxnInfo), ctx)
}