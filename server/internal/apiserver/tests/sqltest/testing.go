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

	statement, err := s.Db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, login VARCHAR(30) NOT NULL UNIQUE, password TEXT NOT NULL)")
	if err != nil {
		t.Fatal(err)
	}
	statement.Exec()
	defer statement.Close()

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.Db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
