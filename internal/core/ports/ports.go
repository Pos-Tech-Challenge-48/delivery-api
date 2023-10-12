package ports

import (
	"context"

	"github.com/ffelipelimao/hex-arch-api/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
}

type UserCreator interface {
	Create(ctx context.Context, user *domain.User) error
}
