package ports

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type ProductRepository interface {
	Add(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, productID string) error
	GetAll(ctx context.Context, params string) ([]domain.Product, error)
	GetCategoryID(ctx context.Context, categoryName string) (categoryID string, err error)
}
