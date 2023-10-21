package ports

import (
	"context"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type CustomerRepository interface {
	Save(ctx context.Context, user *domain.Customer) error
}

type CustomerCreator interface {
	Create(ctx context.Context, user *domain.Customer) error
}
