package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Add(ctx context.Context, product *domain.Product) error {
	query := `
		INSERT INTO product (product_category_id, product_name, product_description, product_unitary_price)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(
		query,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, productID string) error {
	query := `
		DELETE FROM product WHERE product_id = $1
	`
	_, err := r.db.Exec(
		query,
		productID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) GetCategoryID(ctx context.Context, categoryName string) (categoryID string, err error) {

	query := `
		SELECT category_id FROM category WHERE category_name = $1
	`
	row := r.db.QueryRow(query, categoryName)
	err = row.Scan(&categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return categoryID, fmt.Errorf("category not found: %w", err)
		}

		return categoryID, err
	}

	return categoryID, nil
}
