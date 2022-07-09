package model_user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

// User struct ...
type User struct {
	ID                int
	Login             string
	Password          string
	DecryptedPassword string
}

func New() *User {
	return &User{}
}

// Before Create user in database
func (u *User) PreparationCreate() error {
	if len(u.DecryptedPassword) > 0 {
		enc, err := encryptString(u.DecryptedPassword)
		if err != nil {
			return err
		}

		u.Password = enc
	}

	return nil
}

func encryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Validation ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.DecryptedPassword, validation.Required, validation.Length(8, 100)),
	)
}
