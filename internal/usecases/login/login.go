package login

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type LoginUseCase struct {
	customerRepository interfaces.CustomerRepository
}

func NewLoginUseCase(customerRepository interfaces.CustomerRepository) *LoginUseCase {
	return &LoginUseCase{
		customerRepository: customerRepository,
	}
}

func (l *LoginUseCase) SignUp(ctx context.Context, login *entities.Login) (*entities.LoginOutput, error) {
	err := login.Validate()
	if err != nil {
		return nil, err
	}

	customer, err := l.customerRepository.GetByDocumentAndEmail(ctx, login.Document, login.Email)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, entities.ErrInvalidCustomer
	}

	return &entities.LoginOutput{
		JWT: "Bearer jqERnrkenrkebnkafndskandsoanoNSKnskndaknKSAN",
	}, nil
}
