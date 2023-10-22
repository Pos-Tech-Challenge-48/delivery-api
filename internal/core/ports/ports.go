package ports

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
}

type UserCreator interface {
	Create(ctx context.Context, user *domain.User) error
}
