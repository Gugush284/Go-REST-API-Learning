package test_sqlstore

import (
	"testing"

	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
	"github.com/stretchr/testify/assert"
)

func TestImageRepository_Upload(t *testing.T) {
	s, teardown := TestStore(t, dbURL)
	defer teardown("users", "images")

	testcases := []struct {
		image string
		iname string
		txt   string
		name  string
		err   bool
	}{
		{
			name:  "valid",
			image: "filepath",
			txt:   "description",
			iname: "example",
			err:   false,
		},
		{
			name:  "no path",
			image: "",
			txt:   "description",
			iname: "example",
			err:   true,
		},
		{
			name:  "no description",
			image: "filepath",
			txt:   "",
			iname: "example",
			err:   true,
		},
		{
			name:  "bo name",
			image: "filepath",
			txt:   "description",
			iname: "",
			err:   true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			i := &ModelImage.Image{
				Image:     tc.image,
				ImageName: tc.iname,
				Txt:       tc.txt,
			}

			err := s.Image().Upload(i)

			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, i)
			}
		})
	}
}

func TestImageRepository_Download(t *testing.T) {
	s, _ := TestStore(t, dbURL)
	//defer teardown("users", "images")

	i := &ModelImage.Image{
		Image:     "filepath",
		Txt:       "description",
		ImageName: "example",
	}

	err := s.Image().Upload(i)
	assert.NoError(t, err)
	assert.NotNil(t, i)

	im, err := s.Image().Download(i.ImageId)
	assert.NoError(t, err)
	assert.Equal(t, i, im)
}
