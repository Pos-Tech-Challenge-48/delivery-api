package ordersorter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type OrderSorter struct {
	orderRepository interfaces.OrderRepository
}

func NewOrderSorter(orderRepository interfaces.OrderRepository) *OrderSorter {
	return &OrderSorter{
		orderRepository: orderRepository,
	}
}

func (p *OrderSorter) GetAllSortedByStatus(ctx context.Context) ([]entities.Order, error) {
	list, err := p.orderRepository.GetAllSortedByStatus(ctx)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}
