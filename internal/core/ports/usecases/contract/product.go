package contract

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type ProductContract interface {
	Add(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, producID string) error
	GetAll(ctx context.Context, category string) ([]domain.Product, error)
}
