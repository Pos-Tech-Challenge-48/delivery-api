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
		WITH step_one AS (
			INSERT INTO product (product_category_id, product_name, product_description, product_unitary_price)
			VALUES($1, $2, $3, $4)
			RETURNING product_id
		)
		INSERT INTO product_image (product_id, product_image_src_uri)
		SELECT product_id, $5 FROM step_one
	`
	_, err := r.db.Exec(
		query,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Image,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	query := `
		UPDATE product SET product_category_id = $2, product_name = $3, product_description = $4, product_unitary_price = $5
		WHERE product_id = $1
	`
	_, err := r.db.Exec(
		query,
		product.ID,
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

func (r *ProductRepository) UpdateImage(ctx context.Context, productId string, productImage string) error {
	query := `
		UPDATE product_image SET product_image_src_uri = $2
		WHERE product_id = $1
	`
	_, err := r.db.Exec(
		query,
		productId,
		productImage,
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

func (r *ProductRepository) GetAll(ctx context.Context, category string) ([]domain.Product, error) {
	queryParams := []interface{}{}

	query := `
		SELECT
			product.product_id,
			category_name,
			product.product_name,
			product.product_description,
			product_unitary_price,
			COALESCE(product_image.product_image_src_uri, ''),
			product.created_date_db,
			product.last_modified_date_db
		FROM product
        INNER JOIN category
            ON category_id = product_category_id
        LEFT JOIN product_image 
            ON product.product_id = product_image.product_id`

	if category != "" {
		query += ` WHERE category_name = $1`
		queryParams = append(queryParams, category)
	}

	rows, err := r.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.Product, 0)
	var productItem domain.Product

	for rows.Next() {
		err := rows.Scan(
			&productItem.ID,
			&productItem.CategoryID,
			&productItem.Name,
			&productItem.Description,
			&productItem.Price,
			&productItem.Image,
			&productItem.CreatedDate,
			&productItem.LastModifiedDate,
		)
		if err != nil {
			return result, err
		}
		result = append(result, productItem)
	}

	return result, nil

}
func (r *ProductRepository) GetByID(ctx context.Context, ID string) (*domain.Product, error) {
	query := `
		SELECT 
		product_id,
		product_name, 
		product_description, 
		product_unitary_price FROM product WHERE product_id = $1
	`
	row := r.db.QueryRow(query, ID)

	product := &domain.Product{}
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) DeleteImage(ctx context.Context, productID string) error {
	query := `
		DELETE FROM product_image WHERE product_id = $1
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
