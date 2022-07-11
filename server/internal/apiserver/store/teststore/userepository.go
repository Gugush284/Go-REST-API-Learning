package teststore

import (
	globalErrors "github.com/Gugush284/Go-server.git/internal/apiserver"
	ModelUser "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
)

type UserRepository struct {
	store *TestStore
	users map[string]*ModelUser.User
}

// Create ...
func (r *UserRepository) Create(u *ModelUser.User) (*ModelUser.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.PreparationCreate(); err != nil {
		return nil, err
	}

	r.users[u.Login] = u
	u.ID = len(r.users)

	return u, nil
}

// Find by login
func (r *UserRepository) FindByLogin(login string) (*ModelUser.User, error) {
	u, ok := r.users[login]
	if !ok {
		return nil, globalErrors.ErrRecordNotFound
	}

	return u, nil
}
