package ModelUser

import (
	"golang.org/x/crypto/bcrypt"
)

// User struct ...
type User struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password          string `json:"-"`
	DecryptedPassword string `json:"password,omitempty"`
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

func (u *User) Sanitize() {
	u.DecryptedPassword = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
