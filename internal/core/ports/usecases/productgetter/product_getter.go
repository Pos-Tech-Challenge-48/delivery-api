package productgetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type ProductGetter interface {
	GetAll(ctx context.Context, category string) ([]domain.Product, error)
}
