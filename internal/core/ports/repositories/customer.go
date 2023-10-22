package ports

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

//go:generate mockgen -destination=./../../mocks/customerrepositorymock/customer_repository_mock.go -source=./customer.go -package=customerrepositorymock
type CustomerRepository interface {
	Save(ctx context.Context, user *domain.Customer) error
	GetByDocument(ctx context.Context, document string) (*domain.Customer, error)
}
