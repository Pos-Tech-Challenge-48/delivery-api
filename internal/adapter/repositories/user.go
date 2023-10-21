package repositories

import (
	"context"
	"database/sql"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (ur *CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	query := `
	INSERT INTO customer (customer_id, customer_name, customer_email, customer_document, created_date_db, last_modified_date_db)
	VALUES ($1, $2, $3, $4, $5, $6)
`
	_, err := ur.db.Exec(
		query,
		customer.ID,
		customer.Name,
		customer.Email,
		customer.Document,
		customer.CreatedDate,
		customer.LastModifiedDate,
	)
	return err
}
