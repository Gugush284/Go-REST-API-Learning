package test_store

import (
	"testing"

	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users")

	u := model_user.New()
	u.Login = "Examole.org"
	u.Password = "656565"
	u, err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
