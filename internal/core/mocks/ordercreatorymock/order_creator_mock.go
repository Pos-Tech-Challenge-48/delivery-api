// Code generated by MockGen. DO NOT EDIT.
// Source: ./order_creator.go
//
// Generated by this command:
//
//	mockgen -destination=./../../../mocks/ordercreatorymock/order_creator_mock.go -source=./order_creator.go -package=ordercreatorymock
//
// Package ordercreatorymock is a generated GoMock package.
package ordercreatorymock

import (
	context "context"
	reflect "reflect"

	domain "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderCreator is a mock of OrderCreator interface.
type MockOrderCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOrderCreatorMockRecorder
}

// MockOrderCreatorMockRecorder is the mock recorder for MockOrderCreator.
type MockOrderCreatorMockRecorder struct {
	mock *MockOrderCreator
}

// NewMockOrderCreator creates a new mock instance.
func NewMockOrderCreator(ctrl *gomock.Controller) *MockOrderCreator {
	mock := &MockOrderCreator{ctrl: ctrl}
	mock.recorder = &MockOrderCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderCreator) EXPECT() *MockOrderCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderCreator) Create(ctx context.Context, order *domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrderCreatorMockRecorder) Create(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderCreator)(nil).Create), ctx, order)
}
