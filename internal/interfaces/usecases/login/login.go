package login

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

type LoginUseCase interface {
	SignUp(ctx context.Context, customerInput *entities.Login) (*entities.LoginOutput, error)
}
