package repositories_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/external/repositories"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupSuite(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS customer
	(
		customer_id varchar(255) primary key,
		customer_document varchar(14),
		customer_name varchar(150),
		customer_email varchar(80),
		created_date_db timestamp default current_timestamp,
		last_modified_date_db timestamp default current_timestamp
	)
	`)
	assert.Nil(t, err)

	return db
}

func tearDownSuite(db *sql.DB) {
	db.Close()
	db.Exec("drop table customer")
}

func Test_CustomerRepository_GetByDocument(t *testing.T) {
	db := setupSuite(t)

	query := `
	INSERT INTO customer (customer_id, customer_name, customer_email, customer_document, created_date_db, last_modified_date_db)
	VALUES ($1, $2, $3, $4, $5, $6)
`
	_, err := db.Exec(
		query,
		"fake-id",
		"fake-name",
		"fake-email",
		"fake-document",
		time.Now(),
		time.Now(),
	)
	assert.Nil(t, err)

	defer tearDownSuite(db)

	tests := []struct {
		name          string
		customer      *entities.Customer
		expectedError error
	}{
		{
			name: "Should get by document",
			customer: &entities.Customer{
				ID:       "fake-id",
				Name:     "fake-name",
				Email:    "fake-email",
				Document: "fake-document",
			},
			expectedError: nil,
		},
		{
			name: "Should return not found",
			customer: &entities.Customer{
				ID:       "fake-id",
				Name:     "fake-name",
				Email:    "fake-email",
				Document: "fake-document-2",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.NewCustomerRepository(db)
			customer, err := repo.GetByDocument(context.Background(), tt.customer.Document)

			assert.Equal(t, tt.expectedError, err)
			if customer != nil {
				assert.Equal(t, tt.customer.ID, customer.ID)
				assert.Equal(t, tt.customer.Name, customer.Name)
				assert.Equal(t, tt.customer.Email, customer.Email)
				assert.Equal(t, tt.customer.Document, customer.Document)
			}

		})
	}
}

func Test_CustomerRepository_Save(t *testing.T) {
	db := setupSuite(t)

	defer tearDownSuite(db)

	tests := []struct {
		name          string
		customer      *entities.Customer
		expectedError error
	}{
		{
			name: "Should insert a new customer",
			customer: &entities.Customer{
				ID:               uuid.NewString(),
				Name:             "John Doe",
				Email:            "mock@mock.com",
				Document:         "123456789",
				CreatedDate:      time.Now(),
				LastModifiedDate: time.Now(),
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositories.NewCustomerRepository(db)
			err := repo.Save(context.Background(), tt.customer)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
