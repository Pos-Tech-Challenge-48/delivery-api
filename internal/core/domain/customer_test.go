package domain_test

import (
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	name := "John Doe"
	email := "johndoe@example.com"
	document := "12345678901"

	customer := domain.NewCustomer(name, email, document)

	assert.Equal(t, name, customer.Name)
	assert.Equal(t, email, customer.Email)
	assert.Equal(t, document, customer.Document)
}

func TestCustomer_Validate(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		document string
		expected error
	}{
		{"Valid Customer", "johndoe@example.com", "03271168032", nil},
		{"", "johndoe@example.com", "03271168032", domain.ErrCustomerEmptyName},
		{"Empty Email", "", "03271168032", domain.ErrCustomerEmptyEmail},
		{"Invalid Email", "johndoeexample.com", "12345678901", domain.ErrCustomerInvalidEmail},
		{"Invalid Document", "johndoe@example.com", "123456", domain.ErrCustomerInvalidDocument},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer := domain.NewCustomer(tt.name, tt.email, tt.document)
			err := customer.Validate()
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestCustomer_ValidateDocument(t *testing.T) {
	tests := []struct {
		document string
		expected error
	}{
		{"03271168032", nil},
		{"", domain.ErrCustomerEmptyDocument},
		{"123456", domain.ErrCustomerInvalidDocument},
	}

	for _, tt := range tests {
		t.Run(tt.document, func(t *testing.T) {
			customer := domain.NewCustomer("John Doe", "johndoe@example.com", tt.document)
			err := customer.ValidateDocument()
			assert.Equal(t, tt.expected, err)
		})
	}
}
