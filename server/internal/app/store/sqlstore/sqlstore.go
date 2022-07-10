package sqlstore

import (
	"database/sql"

	"github.com/Gugush284/Go-server.git/internal/app/store"
	_ "github.com/go-sql-driver/mysql" // sql driver ...
)

// Store for clients ...
type SqlStore struct {
	DbURL          string
	Db             *sql.DB
	userRepository *UserRepository
}

// New Store ...
func New(URL string) *SqlStore {
	return &SqlStore{
		DbURL: URL,
	}
}

// Open db ...
func (s *SqlStore) Open() error {
	db, err := sql.Open("mysql", s.DbURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Db = db

	return nil
}

// Access for User
func (s *SqlStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Close connection ...
func (s *SqlStore) Close() {
	s.Db.Close()
}

func (s *SqlStore) CreateTables() error {
	if err := s.Open(); err != nil {
		return err
	}

	statement, err := s.Db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, login VARCHAR(30) NOT NULL UNIQUE, password TEXT NOT NULL)")
	if err != nil {
		return err
	}
	statement.Exec()
	defer statement.Close()

	return nil
}
