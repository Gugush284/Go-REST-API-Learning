package model_user

import (
	"testing"
)

// Test user config ...
func TestUser(t *testing.T) *User {
	return &User{
		Login:             "example",
		DecryptedPassword: "ex_password",
	}
}
