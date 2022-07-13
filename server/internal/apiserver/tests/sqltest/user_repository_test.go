package test_sqlstore

import (
	"testing"

	model_user "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
	ModelUserTest "github.com/Gugush284/Go-server.git/internal/apiserver/tests/model/users"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users", "images")

	u, err := s.User().Create(ModelUserTest.TestUser(t))
	/*if u != nil {
		fmt.Println(u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users", "images")

	user := ModelUserTest.TestUser(t)
	/*if user != nil {
		fmt.Println("1", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/

	u, err := s.User().FindByLogin(user.Login)
	/*if u != nil {
		fmt.Println("2", u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/
	assert.Error(t, err)
	assert.Nil(t, u)

	u, err = s.User().Create(&model_user.User{
		Login:             user.Login,
		DecryptedPassword: user.DecryptedPassword,
	})
	/*if u != nil {
		fmt.Println("3", u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, u)

	username, err := s.User().FindByLogin(user.Login)
	/*if username != nil {
		fmt.Println("4", username.ID, username.Login, username.DecryptedPassword, username.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, username)
}

func TestUserRepository_Find(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users", "images")

	u := ModelUserTest.TestUser(t)
	/*if user != nil {
		fmt.Println("1", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/

	u, err := s.User().Create(&model_user.User{
		Login:             u.Login,
		DecryptedPassword: u.DecryptedPassword,
	})
	/*if u != nil {
		fmt.Println("2", u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, u)

	username, err := s.User().Find(u.ID)
	/*if username != nil {
		fmt.Println("3", username.ID, username.Login, username.DecryptedPassword, username.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, username)
}
