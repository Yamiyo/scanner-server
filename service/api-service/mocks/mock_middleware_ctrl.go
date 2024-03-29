// Code generated by MockGen. DO NOT EDIT.
// Source: middleware_response.go

// Package mock_restctl is a generated GoMock package.
package mock_restctl

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockResponseMiddlewareInterface is a mock of ResponseMiddlewareInterface interface.
type MockResponseMiddlewareInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResponseMiddlewareInterfaceMockRecorder
}

// MockResponseMiddlewareInterfaceMockRecorder is the mock recorder for MockResponseMiddlewareInterface.
type MockResponseMiddlewareInterfaceMockRecorder struct {
	mock *MockResponseMiddlewareInterface
}

// NewMockResponseMiddlewareInterface creates a new mock instance.
func NewMockResponseMiddlewareInterface(ctrl *gomock.Controller) *MockResponseMiddlewareInterface {
	mock := &MockResponseMiddlewareInterface{ctrl: ctrl}
	mock.recorder = &MockResponseMiddlewareInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponseMiddlewareInterface) EXPECT() *MockResponseMiddlewareInterfaceMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockResponseMiddlewareInterface) Handle(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", ctx)
}

// Handle indicates an expected call of Handle.
func (mr *MockResponseMiddlewareInterfaceMockRecorder) Handle(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockResponseMiddlewareInterface)(nil).Handle), ctx)
}
