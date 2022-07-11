package ModelUserTest

import (
	"testing"

	ModelUser "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
)

// Test user config ...
func TestUser(t *testing.T) *ModelUser.User {
	return &ModelUser.User{
		Login:             "example",
		DecryptedPassword: "ex_password",
	}
}
