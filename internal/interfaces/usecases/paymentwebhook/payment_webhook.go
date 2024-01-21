package paymentwebhook

import (
	"context"
)

type PaymentWebhook interface {
	Update(ctx context.Context, orderID string) error
}
