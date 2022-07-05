package test_store

import (
	"os"
	"testing"
)

var dbURL string

func TestMain(m *testing.M) {
	dbURL = os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "root:root@tcp(127.0.0.1:3306)/apiserver_test"
	}

	os.Exit(m.Run())
}
