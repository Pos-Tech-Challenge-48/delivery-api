package usecases

import (
	"context"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/ports"
)

type CustomerCreator struct {
	customerRepository ports.CustomerRepository
}

func NewCustomerCreator(customerRepository ports.CustomerRepository) *CustomerCreator {
	return &CustomerCreator{
		customerRepository: customerRepository,
	}
}

func (uc *CustomerCreator) Create(ctx context.Context, customerInput *domain.Customer) error {

	err := customerInput.Validate()
	if err != nil {
		return err
	}

	customer := domain.NewCustomer(customerInput.Name, customerInput.Email, customerInput.Email)

	return uc.customerRepository.Save(ctx, customer)
}
