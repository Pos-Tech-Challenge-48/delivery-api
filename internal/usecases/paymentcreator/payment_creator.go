package paymentcreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	api_interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/api"
	repo_interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type PaymentCreator struct {
	orderRepository repo_interfaces.OrderRepository
	paymentGateway  api_interfaces.PaymentGatewayAPI
}

func NewPaymentCreator(orderRepository repo_interfaces.OrderRepository, paymentGateway api_interfaces.PaymentGatewayAPI) *PaymentCreator {
	return &PaymentCreator{
		orderRepository: orderRepository,
		paymentGateway:  paymentGateway,
	}
}

func (p *PaymentCreator) Create(ctx context.Context, orderPayment *entities.OrderPayment) (*entities.OrderPaymentResponse, error) {
	order, err := p.orderRepository.GetByID(ctx, orderPayment.OrderID)
	if err != nil {
		return nil, err
	}

	err = p.validateOrder(order)
	if err != nil {
		return nil, err
	}

	payment := entities.NewPayment(orderPayment, order)

	paymentResponse, err := p.paymentGateway.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	err = p.orderRepository.Update(ctx, order)
	if err != nil {
		return nil, err
	}

	return &entities.OrderPaymentResponse{
		PaymentID: paymentResponse.ID,
		QRCode:    paymentResponse.QRData,
	}, nil
}

func (p *PaymentCreator) validateOrder(order *entities.Order) error {
	if order == nil {
		return entities.ErrOrderNotExist
	}
	if !order.IsReadyToPayment() {
		return entities.ErrOrderNotReady
	}
	return nil
}
