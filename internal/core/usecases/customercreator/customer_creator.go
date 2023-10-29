package customercreator

import (
	"context"
	"errors"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	ports "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports/repositories"
)

var (
	ErrAlreadyExistsCustomer = errors.New("already exists customer with this document")
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

	existsCustomer, err := uc.customerRepository.GetByDocument(ctx, customerInput.Document)
	if err != nil {
		return err
	}

	if existsCustomer != nil {
		return ErrAlreadyExistsCustomer
	}

	customer := domain.NewCustomer(customerInput.Name, customerInput.Email, customerInput.Document)

	return uc.customerRepository.Save(ctx, customer)
}
