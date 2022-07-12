package test_sqlstore

import (
	"testing"

	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
	"github.com/stretchr/testify/assert"
)

func TestImageRepository_Upload(t *testing.T) {
	s, _ := TestStore(t, dbURL)
	//defer teardown("users")

	i := &ModelImage.Image{
		ImageType: "png",
		Image:     "filepath",
		ImageName: "example",
		Txt:       "description",
	}

	err := s.Image().Upload(i)

	assert.NoError(t, err)
	assert.NotNil(t, i)
}
