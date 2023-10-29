package ordercreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
)

type OrderCreator struct {
	orderRepository   ports.OrderRepository
	productRepository ports.ProductRepository
}

func NewOrderCreator(orderRepository ports.OrderRepository, productRepository ports.ProductRepository) *OrderCreator {
	return &OrderCreator{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (uc *OrderCreator) Create(ctx context.Context, orderInput *domain.Order) error {
	err := orderInput.Validate()
	if err != nil {
		return err
	}

	order := domain.NewOrder(orderInput.CustomerID, orderInput.OrderProduct)

	amount, err := uc.calculateAmount(ctx, order.OrderProduct)
	if err != nil {
		return err
	}

	order.Amount = amount

	return uc.orderRepository.Save(ctx, order)
}

func (uc *OrderCreator) calculateAmount(ctx context.Context, products []domain.OrderProduct) (float64, error) {
	var amount float64

	for _, p := range products {

		product, err := uc.productRepository.GetByID(ctx, p.ID)
		if err != nil {
			return 0, err
		}
		amount += product.Price

	}
	return amount, nil

}
