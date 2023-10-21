package domain

import (
	"errors"
	"time"
)

type Customer struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Document         string    `json:"document"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
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
