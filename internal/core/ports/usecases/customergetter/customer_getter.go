package customergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

//go:generate mockgen -destination=./../../../mocks/customergetterymock/customer_getter_mock.go -source=./customer_getter.go -package=customergetterymock
type CustomerGetter interface {
	Get(ctx context.Context, customerInput *domain.Customer) (*domain.Customer, error)
}
