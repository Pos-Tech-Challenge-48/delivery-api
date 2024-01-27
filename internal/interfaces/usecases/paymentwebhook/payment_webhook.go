package paymentwebhook

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

type PaymentWebhook interface {
	Update(ctx context.Context, orderID string, paymentRequest entities.OrderPaymentRequest) error
}
