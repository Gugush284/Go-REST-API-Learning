package test_store

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Gugush284/Go-server.git/internal/app/store"
)

// TEST STORE ...
func TestStore(t *testing.T, DbURL string) (*store.Store, func(...string)) {
	t.Helper()

	config := store.NewConfig()
	config.DatabaseURL = DbURL
	s := store.New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	db := s.Return_connection()
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, login VARCHAR(30) NOT NULL UNIQUE, password TEXT NOT NULL)")
	if err != nil {
		t.Fatal(err)
	}
	statement.Exec()
	defer statement.Close()

	return s, func(tables ...string) {
		if len(tables) > 0 {
			db := s.Return_connection()
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
