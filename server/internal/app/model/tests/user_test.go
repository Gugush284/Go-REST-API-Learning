package tests_model

import (
	"testing"

	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
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
		u       func() *model_user.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model_user.User {
				return TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty login",
			u: func() *model_user.User {
				u := TestUser(t)
				u.Login = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not size login",
			u: func() *model_user.User {
				u := TestUser(t)
				u.Login = "ww"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model_user.User {
				u := TestUser(t)
				u.DecryptedPassword = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model_user.User {
				u := TestUser(t)
				u.DecryptedPassword = "no"
				return u
			},
			isValid: false,
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
