package ModelImageTest

import (
	"testing"

	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
)

func TestImage(t *testing.T) *ModelImage.Image {
	return &ModelImage.Image{
		Image:     "filepath",
		ImageName: "example",
		Txt:       "-",
	}
}
