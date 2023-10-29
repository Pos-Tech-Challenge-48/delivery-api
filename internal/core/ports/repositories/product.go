package ports

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type ProductRepository interface {
	Add(ctx context.Context, user *domain.Product) error
	GetCategoryID(ctx context.Context, categoryName string) (categoryID string, err error)
}
