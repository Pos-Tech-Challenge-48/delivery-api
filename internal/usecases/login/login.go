package login

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfacesrapi "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/api"
	interfacesrepo "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/repositories"
)

type LoginUseCase struct {
	customerRepository interfacesrepo.CustomerRepository
	authAPI            interfacesrapi.AuthorizerAPI
}

func NewLoginUseCase(customerRepository interfacesrepo.CustomerRepository, authAPI interfacesrapi.AuthorizerAPI) *LoginUseCase {
	return &LoginUseCase{
		customerRepository: customerRepository,
		authAPI:            authAPI,
	}
}

func (l *LoginUseCase) Login(ctx context.Context, login *entities.Login) (*entities.LoginOutput, error) {
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

	token, err := l.authAPI.Execute(ctx, login)
	if err != nil {
		return nil, err
	}

	return &entities.LoginOutput{
		JWT: token,
	}, nil
}
