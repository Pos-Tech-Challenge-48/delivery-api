package ordergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type OrderGetter interface {
	GetAll(ctx context.Context) ([]domain.Order, error)
}
