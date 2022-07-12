package store

import (
	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
	model_user "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
)

type UserRepository interface {
	Create(*model_user.User) (*model_user.User, error)
	FindByLogin(string) (*model_user.User, error)
	Find(int) (*model_user.User, error)
}

type ImageRepository interface {
	Upload(i *ModelImage.Image) error
}
