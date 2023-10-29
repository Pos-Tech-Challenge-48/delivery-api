package ports

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

//go:generate mockgen -destination=./../../mocks/orderrepositorymock/order_repository_mock.go -source=./order.go -package=orderrepositorymock
type OrderRepository interface {
	Save(ctx context.Context, order *domain.Order) error
}
