package api

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
)

type AuthorizerAPI interface {
	Execute(ctx context.Context, login *entities.Login) (string, error)
}
