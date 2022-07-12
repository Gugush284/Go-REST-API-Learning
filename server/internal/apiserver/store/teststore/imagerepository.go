package teststore

import ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"

type ImageRepository struct {
	store    *TestStore
	imagesID map[int]*ModelImage.Image
}

func (r *ImageRepository) Upload(i *ModelImage.Image) error {
	if err := i.Validate(); err != nil {
		return err
	}

	i.ImageId = len(r.imagesID) + 1
	r.imagesID[i.ImageId] = i

	return nil
}
