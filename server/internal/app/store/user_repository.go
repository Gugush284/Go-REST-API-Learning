package store

import (
	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model_user.User) (*model_user.User, error) {

	if err := r.store.db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	statement, err := r.store.db.Prepare("INSERT INTO users (login, password) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	statement.Exec(u.Login, u.Password)

	row := r.store.db.QueryRow("SELECT LAST_INSERT_ID()")

	if err := row.Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// Find by login
func (r *UserRepository) FindByLogin(login string) (*model_user.User, error) {
	if err := r.store.db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	u := model_user.New()
	u.Login = login

	row := r.store.db.QueryRow(
		"SELECT id, password FROM users WHERE login = (?)",
		login)
	if err := row.Scan(&u.ID, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
