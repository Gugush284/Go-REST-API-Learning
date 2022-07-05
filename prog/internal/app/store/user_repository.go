package store

import models "github.com/Gugush284/Go-server.git/internal/app/model/user"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	if err := r.store.Open(); err != nil {
		return nil, err
	}

	statement, err := r.store.db.Prepare("INSERT INTO users (login, password) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	statement.Exec(u.GetLogin(), u.GetPassword())

	rows, err := r.store.db.Query("SELECT LAST_INSERT_ID()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	rows.Scan(&id)
	u.Add_id(id)

	return u, nil
}
