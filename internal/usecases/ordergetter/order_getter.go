package ordergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type OrderGetter struct {
	orderRepository   interfaces.OrderRepository
	productRepository interfaces.ProductRepository
}

func NewOrderGetter(orderRepository interfaces.OrderRepository, productRepository interfaces.ProductRepository) *OrderGetter {
	return &OrderGetter{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (p *OrderGetter) GetAll(ctx context.Context) ([]entities.Order, error) {

	list, err := p.orderRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}