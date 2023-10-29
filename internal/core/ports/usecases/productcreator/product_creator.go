package productcreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type ProductCreator interface {
	Add(ctx context.Context, product *domain.Product) error
}
