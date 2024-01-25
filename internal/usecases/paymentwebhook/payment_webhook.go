package paymentwebhook

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	repo_interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type PaymentWebhook struct {
	orderRepository repo_interfaces.OrderRepository
}

func NewPaymentWebhook(orderRepository repo_interfaces.OrderRepository) *PaymentWebhook {
	return &PaymentWebhook{
		orderRepository: orderRepository,
	}
}

func (p *PaymentWebhook) Update(ctx context.Context, orderID string) error {
	order, err := p.orderRepository.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	err = p.validateOrder(order)
	if err != nil {
		return err
	}

	order = order.SetPaid()

	err = p.orderRepository.Update(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentWebhook) validateOrder(order *entities.Order) error {
	if order == nil {
		return entities.ErrOrderNotExist
	}
	if !order.IsReadyToConfirmPayment() {
		return entities.ErrOrderNotReadyToConfirmPayment
	}
	return nil
}
