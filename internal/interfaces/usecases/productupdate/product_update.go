package productupdate

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

type ProductUpdate interface {
	Update(ctx context.Context, data *entities.Product) error
}
