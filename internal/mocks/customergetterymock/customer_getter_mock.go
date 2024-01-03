// Code generated by MockGen. DO NOT EDIT.
// Source: ./customer_getter.go
//
// Generated by this command:
//
//	mockgen -destination=./../../../mocks/customergetterymock/customer_getter_mock.go -source=./customer_getter.go -package=customergetterymock
//
// Package customergetterymock is a generated GoMock package.
package customergetterymock

import (
	context "context"
	reflect "reflect"

	entities "github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockCustomerGetter is a mock of CustomerGetter interface.
type MockCustomerGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerGetterMockRecorder
}

// MockCustomerGetterMockRecorder is the mock recorder for MockCustomerGetter.
type MockCustomerGetterMockRecorder struct {
	mock *MockCustomerGetter
}

// NewMockCustomerGetter creates a new mock instance.
func NewMockCustomerGetter(ctrl *gomock.Controller) *MockCustomerGetter {
	mock := &MockCustomerGetter{ctrl: ctrl}
	mock.recorder = &MockCustomerGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerGetter) EXPECT() *MockCustomerGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCustomerGetter) Get(ctx context.Context, customerInput *entities.Customer) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, customerInput)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCustomerGetterMockRecorder) Get(ctx, customerInput any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCustomerGetter)(nil).Get), ctx, customerInput)
}