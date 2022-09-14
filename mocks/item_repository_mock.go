// Code generated by MockGen. DO NOT EDIT.
// Source: ./item_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockItemRepository is a mock of ItemRepository interface
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// SaveItem mocks base method
func (m *MockItemRepository) SaveItem(name string, stock int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveItem", name, stock)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveItem indicates an expected call of SaveItem
func (mr *MockItemRepositoryMockRecorder) SaveItem(name, stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveItem", reflect.TypeOf((*MockItemRepository)(nil).SaveItem), name, stock)
}

// GetItemByID mocks base method
func (m *MockItemRepository) GetItemByID(itemID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemByID", itemID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetItemByID indicates an expected call of GetItemByID
func (mr *MockItemRepositoryMockRecorder) GetItemByID(itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemByID", reflect.TypeOf((*MockItemRepository)(nil).GetItemByID), itemID)
}
