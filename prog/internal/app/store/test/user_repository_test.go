package test_store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users")

	u, err := s.User().Create("Examole.org", "656565")

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
