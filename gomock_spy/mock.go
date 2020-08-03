// Code generated by MockGen. DO NOT EDIT.
// Source: DeleteCronTask.go

// Package gomock_spy is a generated GoMock package.
package gomock_spy

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeleteService is a mock of DeleteService interface
type MockDeleteService struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteServiceMockRecorder
}

// MockDeleteServiceMockRecorder is the mock recorder for MockDeleteService
type MockDeleteServiceMockRecorder struct {
	mock *MockDeleteService
}

// NewMockDeleteService creates a new mock instance
func NewMockDeleteService(ctrl *gomock.Controller) *MockDeleteService {
	mock := &MockDeleteService{ctrl: ctrl}
	mock.recorder = &MockDeleteServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeleteService) EXPECT() *MockDeleteServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockDeleteService) Delete(indexName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", indexName)
}

// Delete indicates an expected call of Delete
func (mr *MockDeleteServiceMockRecorder) Delete(indexName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeleteService)(nil).Delete), indexName)
}
