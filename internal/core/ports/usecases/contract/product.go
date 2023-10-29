package contract

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type ProductContract interface {
	Add(ctx context.Context, user *domain.Product) error
}
