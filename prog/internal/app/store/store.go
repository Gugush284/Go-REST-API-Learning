package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // sql driver ...
)

// Store for clients ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New Store ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open db ...
func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Return_connection() *sql.DB {
	return s.db
}

// Access for User
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return &UserRepository{}
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Close connection ...
func (s *Store) Close() {
	s.db.Close()
}
