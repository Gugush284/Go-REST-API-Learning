package apiserver

import (
	"net/http"

	"github.com/Gugush284/Go-server.git/internal/apiserver/store/sqlstore"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	store := sqlstore.New(config.DatabaseURL)

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	srv := NewServer(store, sessionStore)

	if err := srv.configureLogger(config); err != nil {
		return err
	}

	if err := store.CreateTables(); err != nil {
		return err
	}

	srv.Logger.Info("starting api server")
	srv.Logger.Debug(config.SessionKey)
	defer store.Db.Close()

	return http.ListenAndServe(config.BindAddr, srv)
}
