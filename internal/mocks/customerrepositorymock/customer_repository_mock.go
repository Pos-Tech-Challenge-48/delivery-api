// Code generated by MockGen. DO NOT EDIT.
// Source: ./customer.go
//
// Generated by this command:
//
//	mockgen -destination=./../../mocks/customerrepositorymock/customer_repository_mock.go -source=./customer.go -package=customerrepositorymock
//
// Package customerrepositorymock is a generated GoMock package.
package customerrepositorymock

import (
	context "context"
	reflect "reflect"

	entities "github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// GetByDocument mocks base method.
func (m *MockCustomerRepository) GetByDocument(ctx context.Context, document string) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDocument", ctx, document)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDocument indicates an expected call of GetByDocument.
func (mr *MockCustomerRepositoryMockRecorder) GetByDocument(ctx, document any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDocument", reflect.TypeOf((*MockCustomerRepository)(nil).GetByDocument), ctx, document)
}

// Save mocks base method.
func (m *MockCustomerRepository) Save(ctx context.Context, user *entities.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockCustomerRepositoryMockRecorder) Save(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCustomerRepository)(nil).Save), ctx, user)
}