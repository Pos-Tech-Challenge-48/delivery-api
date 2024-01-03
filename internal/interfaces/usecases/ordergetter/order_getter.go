package ordergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

type OrderGetter interface {
	GetAll(ctx context.Context) ([]entities.Order, error)
}
