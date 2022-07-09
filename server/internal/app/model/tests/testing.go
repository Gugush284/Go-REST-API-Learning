package tests_model

import (
	"testing"

	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
)

// Test user config ...
func TestUser(t *testing.T) *model_user.User {
	return &model_user.User{
		Login:             "example",
		DecryptedPassword: "ex_password",
	}
}
