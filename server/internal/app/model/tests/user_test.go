package tests_model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_PreparationCreate(t *testing.T) {
	u := TestUser(t)
	assert.NoError(t, u.PreparationCreate())
	assert.NotNil(t, u.Password)
}
