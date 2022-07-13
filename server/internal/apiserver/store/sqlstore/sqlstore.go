package sqlstore

import (
	"database/sql"

	"github.com/Gugush284/Go-server.git/internal/apiserver/store"
	_ "github.com/go-sql-driver/mysql" // sql driver ...
)

// Store for clients ...
type SqlStore struct {
	DbURL           string
	Db              *sql.DB
	userRepository  *UserRepository
	imageRepository *ImageRepository
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

// Access for Image
func (s *SqlStore) Image() store.ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = &ImageRepository{
		store: s,
	}

	return s.imageRepository
}

// Close connection ...
func (s *SqlStore) Close() {
	s.Db.Close()
}

func (s *SqlStore) CreateTables() error {
	if err := s.Open(); err != nil {
		return err
	}

	statement, err := s.Db.Prepare(`create table IF NOT EXISTS users (
		id integer not null PRIMARY KEY AUTO_INCREMENT,
		login varchar(30) not null UNIQUE,
		password TEXT not null
	)`)
	if err != nil {
		return err
	}
	statement.Exec()
	statement.Close()

	statement, err = s.Db.Prepare(`create table IF NOT EXISTS images (
		image_id        integer     not null PRIMARY KEY AUTO_INCREMENT,
		image           varchar(50) not null default '',
		image_name      varchar(50) not null default '',
		txt				Text		not null
	)`)
	if err != nil {
		return err
	}
	statement.Exec()
	statement.Close()

	return nil
}
