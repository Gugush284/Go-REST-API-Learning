package teststore

import (
	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
	model_user "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
	"github.com/Gugush284/Go-server.git/internal/apiserver/store"
)

// Store for clients ...
type TestStore struct {
	userRepository  *UserRepository
	imageRepository *ImageRepository
}

// New Store ...
func New() *TestStore {
	return &TestStore{}
}

// Access for User
func (s *TestStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store:    s,
		usersStr: make(map[string]*model_user.User),
		usersID:  make(map[int]*model_user.User),
	}

	return s.userRepository
}

// Access for Image
func (s *TestStore) Image() store.ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = &ImageRepository{
		store:    s,
		imagesID: make(map[int]*ModelImage.Image),
	}

	return s.imageRepository
}
