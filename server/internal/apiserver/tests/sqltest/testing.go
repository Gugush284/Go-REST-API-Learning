package test_sqlstore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Gugush284/Go-server.git/internal/apiserver/store/sqlstore"
)

// TEST STORE ...
func TestStore(t *testing.T, DbURL string) (*sqlstore.SqlStore, func(...string)) {
	t.Helper()

	s := sqlstore.New(DbURL)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	statement, err := s.Db.Prepare(`create table IF NOT EXISTS users (
		id integer not null PRIMARY KEY AUTO_INCREMENT,
		login varchar(30) not null UNIQUE,
		password TEXT not null
	)`)
	if err != nil {
		t.Fatal(err)
	}
	statement.Exec()
	statement.Close()

	statement, err = s.Db.Prepare(`create table IF NOT EXISTS images (
		image_id        integer     not null PRIMARY KEY AUTO_INCREMENT,
		image_type      varchar(25) not null default '',
		image           varchar(50) not null default '',
		image_name      varchar(50) not null default '',
		txt				Text		not null
	)`)
	if err != nil {
		t.Fatal(err)
	}
	statement.Exec()
	statement.Close()

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.Db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
