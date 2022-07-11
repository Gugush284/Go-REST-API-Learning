package teststore

import (
	globalErrors "github.com/Gugush284/Go-server.git/internal/apiserver"
	ModelUser "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
)

type UserRepository struct {
	store    *TestStore
	usersStr map[string]*ModelUser.User
	usersID  map[int]*ModelUser.User
}

// Create ...
func (r *UserRepository) Create(u *ModelUser.User) (*ModelUser.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.PreparationCreate(); err != nil {
		return nil, err
	}

	r.usersStr[u.Login] = u
	u.ID = len(r.usersStr)
	r.usersID[u.ID] = u

	return u, nil
}

// Find by login
func (r *UserRepository) FindByLogin(login string) (*ModelUser.User, error) {
	u, ok := r.usersStr[login]
	if !ok {
		return nil, globalErrors.ErrRecordNotFound
	}

	return u, nil
}

// Find by id
func (r *UserRepository) Find(id int) (*ModelUser.User, error) {
	u, ok := r.usersID[id]
	if !ok {
		return nil, globalErrors.ErrRecordNotFound
	}

	return u, nil
}
