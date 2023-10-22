package customergetter

import (
	"context"
	"errors"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
)

var (
	ErrCustomerNotFound = errors.New("customer not find with this document")
)

type CustomerGetter struct {
	customerRepository ports.CustomerRepository
}

func NewCustomerGetter(customerRepository ports.CustomerRepository) *CustomerGetter {
	return &CustomerGetter{
		customerRepository: customerRepository,
	}
}

func (uc *CustomerGetter) Get(ctx context.Context, customerInput *domain.Customer) (*domain.Customer, error) {
	err := customerInput.ValidateDocument()
	if err != nil {
		return nil, err
	}

	customer, err := uc.customerRepository.GetByDocument(ctx, customerInput.Document)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, ErrCustomerNotFound
	}

	return customer, nil
}
