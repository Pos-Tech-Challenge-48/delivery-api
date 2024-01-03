package interfaces

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

//go:generate mockgen -destination=./../../mocks/orderrepositorymock/order_repository_mock.go -source=./order.go -package=orderrepositorymock
type OrderRepository interface {
	Save(ctx context.Context, order *entities.Order) error
	GetAll(ctx context.Context) ([]entities.Order, error)
}
