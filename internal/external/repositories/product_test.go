package repositories_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/external/repositories"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupSuite_p(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS category (
		category_id varchar(255) constraint category_pk primary key,
		category_name varchar(100) not null
	);
	
	CREATE TABLE IF NOT EXISTS product
	(
		product_id varchar(255) primary key,
		product_category_id varchar(255) constraint product_category_id_fk references category (category_id),
		product_name varchar(100) not null,
		product_description varchar,
		product_unitary_price NUMERIC(4, 2) not null,
		created_date_db timestamp default current_timestamp,
		last_modified_date_db timestamp default current_timestamp
	);

	INSERT INTO category(category_id, category_name)
	VALUES
		('c8461b05-a06b-4d63-a05d-0d2f43661985','Lanche'),
		('35d246ae-6120-4b57-87e1-3db9048d5329','Acompanhamento'),
		('f2868a24-7498-47a0-b98b-2690c59bf658','Bebida'),
		('807d7134-e478-40e2-8805-b26836a023d2','Sobremesa');
	`)
	assert.Nil(t, err)

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS product_image
	(
		product_id varchar(255) primary key,
		product_image_src_uri varchar(255)
	);`)
	assert.Nil(t, err)

	return db
}

func tearDownSuite_p(db *sql.DB) {
	db.Close()
	db.Exec("drop table product")
}

func Test_ProductRepository_Add(t *testing.T) {
	db := setupSuite_p(t)

	query := `
	INSERT INTO product (product_id, product_category_id, product_name, product_description, product_unitary_price)
	VALUES ($1, $2, $3, $4, $5)
`
	_, err := db.Exec(
		query,
		"fake-id",
		"fake-category",
		"fake-name",
		"fake-description",
		23.0,
	)
	assert.Nil(t, err)

	defer tearDownSuite_p(db)

	tests := []struct {
		name          string
		product       *entities.Product
		expectedError error
	}{
		{
			name: "Inserting product successful",
			product: &entities.Product{
				CategoryID:  "fake-cat-id",
				Name:        "fake-name",
				Description: "fake-description",
				Image:       "fake-image",
				Price:       9,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.NewProductRepository(db)
			err := repo.Add(context.Background(), tt.product)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
