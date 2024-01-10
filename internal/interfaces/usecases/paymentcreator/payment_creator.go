package paymentcreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

type PaymentCreator interface {
	Create(ctx context.Context, orderPayment *entities.OrderPayment) (*entities.OrderPaymentResponse, error)
}
