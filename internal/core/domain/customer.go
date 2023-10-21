package domain

import (
	"errors"
	"math/rand"
	"time"
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
		return errors.New("empty name")
	}

	if u.Email == "" {
		return errors.New("empty email")
	}

	return nil
}

func (u *Customer) ValidateDocument() error {

	if u.Document == "" {
		return errors.New("empty document")
	}

	return nil
}

func (u *Customer) getID() int64 {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	random := generator.Intn(100) + 1

	return int64(random)
}
