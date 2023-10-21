package usecases

import (
	"context"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/ports"
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

	return customer, nil
}
