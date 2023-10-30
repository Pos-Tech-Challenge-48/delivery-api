package ordergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
	"github.com/pkg/errors"
)

type OrderGetter struct {
	orderRepository   ports.OrderRepository
	productRepository ports.ProductRepository
}

func NewOrderGetter(orderRepository ports.OrderRepository, productRepository ports.ProductRepository) *OrderGetter {
	return &OrderGetter{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (p *OrderGetter) GetAll(ctx context.Context) ([]domain.Order, error) {

	list, err := p.orderRepository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting list of product")
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}
