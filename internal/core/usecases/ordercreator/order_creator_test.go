package ordercreator_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/mocks/orderrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/mocks/productrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/ordercreator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OrderCreator(t *testing.T) {
	tests := []struct {
		name          string
		input         *domain.Order
		expectedError error
		executeMock   func(input *domain.Order, m *orderrepositorymock.MockOrderRepository,
			pm *productrepositorymock.MockProductRepository)
	}{
		{
			name:          "Should create order with success",
			expectedError: nil,
			input: &domain.Order{
				CustomerID: "uuid",
				OrderProduct: []domain.OrderProduct{
					{
						ID:   "uuid",
						Name: "Coca",
					},
					{
						ID:   "uuid2",
						Name: "X-burger",
					},
				},
			},
			executeMock: func(input *domain.Order, m *orderrepositorymock.MockOrderRepository, pm *productrepositorymock.MockProductRepository) {
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
				pm.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&domain.Product{}, nil).MaxTimes(2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			OrderRepositoryMock := orderrepositorymock.NewMockOrderRepository(ctrl)
			productRepositoryMock := productrepositorymock.NewMockProductRepository(ctrl)

			tt.executeMock(tt.input, OrderRepositoryMock, productRepositoryMock)

			OrderCreator := ordercreator.NewOrderCreator(OrderRepositoryMock, productRepositoryMock)

			err := OrderCreator.Create(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
