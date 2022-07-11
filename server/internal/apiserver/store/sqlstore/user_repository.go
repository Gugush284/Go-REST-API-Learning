package sqlstore

import (
	"database/sql"

	Constants "github.com/Gugush284/Go-server.git/internal/apiserver"
	model_user "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
)

// UserRepository ...
type UserRepository struct {
	store *SqlStore
}

// Create ...
func (r *UserRepository) Create(u *model_user.User) (*model_user.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	if err := u.PreparationCreate(); err != nil {
		return nil, err
	}

	statement, err := r.store.Db.Exec("INSERT INTO users (login, password) VALUES (?, ?)", u.Login, u.Password)
	if err != nil {
		return nil, err
	}

	id, err := statement.LastInsertId()
	if err != nil {
		return nil, err
	}
	if id == 0 {
		return nil, nil
	}

	u.ID = int(id)

	return u, nil
}

// Find by login
func (r *UserRepository) FindByLogin(login string) (*model_user.User, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	u := model_user.New()
	u.Login = login

	row := r.store.Db.QueryRow(
		"SELECT id, password FROM users WHERE login = (?)",
		login)
	if err := row.Scan(&u.ID, &u.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, Constants.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Find by id
func (r *UserRepository) Find(id int) (*model_user.User, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	u := model_user.New()
	u.ID = id

	row := r.store.Db.QueryRow(
		"SELECT login, password FROM users WHERE id = (?)",
		id)
	if err := row.Scan(&u.Login, &u.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, Constants.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
