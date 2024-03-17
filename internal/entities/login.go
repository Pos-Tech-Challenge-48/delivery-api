package entities

import (
	"errors"

	"github.com/klassmann/cpfcnpj"
)

var ErrInvalidCustomer = errors.New("email or document is wrong")

type Login struct {
	Email    string `json:"email"`
	Document string `json:"document"`
	Password string `json:"password"`
}

type LoginOutput struct {
	JWT string `json:"jwt"`
}

func (u *Login) Validate() error {
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
