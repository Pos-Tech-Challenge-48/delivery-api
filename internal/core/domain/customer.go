package domain

import (
	"errors"
	"math/rand"
	"regexp"
	"time"

	"github.com/klassmann/cpfcnpj"
)

var (
	ErrCustomerEmptyEmail      = errors.New("customer empty email")
	ErrCustomerInvalidEmail    = errors.New("customer invalid email")
	ErrCustomerEmptyName       = errors.New("customer empty name")
	ErrCustomerEmptyDocument   = errors.New("customer empty document")
	ErrCustomerInvalidDocument = errors.New("customer invalid document")
)

var (
	EmailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

type Customer struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Document         string    `json:"document"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
}

func NewCustomer(name string, email string, document string) *Customer {

	c := &Customer{
		Name:             name,
		Email:            email,
		Document:         document,
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}

	c.ID = c.getID()
	return c
}

func (u *Customer) Validate() error {
	if u.Name == "" {
		return ErrCustomerEmptyName
	}

	if u.Email == "" {
		return ErrCustomerEmptyEmail
	}

	if !EmailRegex.MatchString(u.Email) {
		return ErrCustomerInvalidEmail
	}

	if isValid := cpfcnpj.ValidateCPF(u.Document); !isValid {
		return ErrCustomerInvalidDocument
	}

	return nil
}

func (u *Customer) ValidateDocument() error {

	if u.Document == "" {
		return ErrCustomerEmptyDocument
	}

	if isValid := cpfcnpj.ValidateCPF(u.Document); !isValid {
		return ErrCustomerInvalidDocument
	}

	return nil
}

func (u *Customer) getID() int64 {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	random := generator.Intn(100) + 1

	return int64(random)
}
