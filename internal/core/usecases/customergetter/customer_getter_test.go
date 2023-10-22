package customergetter_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/mocks/customerrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customergetter"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ConsumerGetter(t *testing.T) {
	tests := []struct {
		name          string
		input         *domain.Customer
		output        *domain.Customer
		expectedError error
		executeMock   func(input, output *domain.Customer, m *customerrepositorymock.MockCustomerRepository)
	}{
		{
			name:          "Should get customer with success",
			expectedError: nil,
			input: &domain.Customer{
				Document: "52411797044",
			},
			output: &domain.Customer{
				ID:       "fake-id",
				Name:     "John Doe",
				Email:    "mock@mock.com",
				Document: "52411797044",
			},
			executeMock: func(input, output *domain.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(output, nil)
			},
		},
		{
			name:          "Should return not found",
			expectedError: customergetter.ErrCustomerNotFound,
			input: &domain.Customer{
				Document: "52411797044",
			},
			output: nil,
			executeMock: func(input, output *domain.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(nil, nil)
			},
		},
		{
			name:          "Should not get invalid document",
			expectedError: domain.ErrCustomerInvalidDocument,
			input: &domain.Customer{
				Document: "XXXXXX",
			},
			executeMock: func(input, output *domain.Customer, m *customerrepositorymock.MockCustomerRepository) {},
		},
		{
			name:          "Should expected an error to get customer",
			expectedError: assert.AnError,
			input: &domain.Customer{
				Document: "52411797044",
			},
			executeMock: func(input, output *domain.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(nil, assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepositoryMock := customerrepositorymock.NewMockCustomerRepository(ctrl)

			tt.executeMock(tt.input, tt.output, customerRepositoryMock)

			customerGetter := customergetter.NewCustomerGetter(customerRepositoryMock)

			customer, err := customerGetter.Get(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, customer)
		})
	}
}
