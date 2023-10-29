package productupdate

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type ProductUpdate interface {
	Update(ctx context.Context, data *domain.Product) error
}
