package teststore

import (
	globalErrors "github.com/Gugush284/Go-server.git/internal/app"
	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
)

type UserRepository struct {
	store *TestStore
	users map[string]*model_user.User
}

// Create ...
func (r *UserRepository) Create(u *model_user.User) (*model_user.User, error) {
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
func (r *UserRepository) FindByLogin(login string) (*model_user.User, error) {
	u, ok := r.users[login]
	if !ok {
		return nil, globalErrors.ErrRecordNotFound
	}

	return u, nil
}
