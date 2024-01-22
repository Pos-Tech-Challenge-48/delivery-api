package api

import "github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"

type PaymentGatewayAPI interface {
	CreatePayment(payment *entities.Payment) (*entities.PaymentResponse, error)
}
