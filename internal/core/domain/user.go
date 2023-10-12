package domain

import "errors"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("empty name")
	}

	if u.Email == "" {
		return errors.New("empty email")
	}

	return nil
}
