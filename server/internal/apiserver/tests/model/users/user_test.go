package ModelUserTest

import (
	"testing"

	ModelUser "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
	"github.com/stretchr/testify/assert"
)

func TestUser_PreparationCreate(t *testing.T) {
	u := TestUser(t)
	assert.NoError(t, u.PreparationCreate())
	assert.NotNil(t, u.Password)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *ModelUser.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *ModelUser.User {
				return TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty login",
			u: func() *ModelUser.User {
				u := TestUser(t)
				u.Login = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not size login",
			u: func() *ModelUser.User {
				u := TestUser(t)
				u.Login = "ww"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *ModelUser.User {
				u := TestUser(t)
				u.DecryptedPassword = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *ModelUser.User {
				u := TestUser(t)
				u.DecryptedPassword = "no"
				return u
			},
			isValid: false,
		},
		{
			name: "with encrypted password",
			u: func() *ModelUser.User {
				u := TestUser(t)
				u.DecryptedPassword = ""
				u.Password = "%asdw%5656"
				return u
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				if tc.isValid {
					assert.NoError(t, tc.u().Validate())
				} else {
					assert.Error(t, tc.u().Validate())
				}
			},
		)
	}
}
