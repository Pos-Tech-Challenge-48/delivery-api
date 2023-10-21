package repositories

import (
	"context"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (ur *CustomerRepository) Save(ctx context.Context, Customer *domain.Customer) error {
	return nil
}
