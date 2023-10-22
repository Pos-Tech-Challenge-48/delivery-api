package customercreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

//go:generate mockgen -destination=./../../../mocks/customercreatorymock/customer_creator_mock.go -source=./customer_creator.go -package=customercreatorymock
type CustomerCreator interface {
	Create(ctx context.Context, user *domain.Customer) error
}
