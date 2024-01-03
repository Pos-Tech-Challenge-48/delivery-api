package productgetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

type ProductGetter interface {
	GetAll(ctx context.Context, category string) ([]entities.Product, error)
}
