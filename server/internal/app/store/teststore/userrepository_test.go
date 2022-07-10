package teststore

import (
	"testing"

	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := New()
	u, err := s.User().Create(model_user.TestUser(t))
	/*if u != nil {
		fmt.Println(u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := New()
	u := model_user.TestUser(t)
	/*if u != nil {
		fmt.Println("1", u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/

	user, err := s.User().FindByLogin(u.Login)
	/*if user != nil {
		fmt.Println("2", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = s.User().Create(&model_user.User{
		Login:             u.Login,
		DecryptedPassword: u.DecryptedPassword,
	})
	/*if user != nil {
		fmt.Println("3", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, user)

	user, err = s.User().FindByLogin(u.Login)
	/*if user != nil {
		fmt.Println("4", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
