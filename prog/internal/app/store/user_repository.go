package store

import (
	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(login string, password string) (*model_user.User, error) {
	if err := r.store.Open(); err != nil {
		return nil, err
	}

	statement, err := r.store.db.Prepare("INSERT INTO users (login, password) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	statement.Exec(login, password)

	rows, err := r.store.db.Query("SELECT LAST_INSERT_ID()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	rows.Scan(&id)

	u := model_user.New()
	u.Add_User(id, login, password)

	return u, nil
}
