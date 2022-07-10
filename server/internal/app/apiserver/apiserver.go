package apiserver

import (
	"net/http"

	"github.com/Gugush284/Go-server.git/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	store := sqlstore.New(config.DatabaseURL)
	srv := newServer(store)

	if err := srv.configureLogger(config); err != nil {
		return err
	}

	if err := store.CreateTables(); err != nil {
		return err
	}

	srv.logger.Info("starting api server")
	defer store.Db.Close()

	return http.ListenAndServe(config.BindAddr, srv)
}
