package test_store

import (
	"fmt"
	"testing"

	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users")

	u := model_user.New()
	u.Login = "Examale.org"
	u.Password = "6565"
	u, err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users")

	login := "username@example.ex"

	u, err := s.User().FindByLogin(login)
	if u != nil {
		fmt.Println(u.ID, u.Login, u.Password)
	}
	assert.Error(t, err)
	assert.Nil(t, u)

	u, err = s.User().Create(&model_user.User{
		Login:    login,
		Password: "1234",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)

	username, err := s.User().FindByLogin(login)
	if u != nil {
		fmt.Println(u.ID, u.Login, u.Password)
	}
	assert.NoError(t, err)
	assert.NotNil(t, username)
}
