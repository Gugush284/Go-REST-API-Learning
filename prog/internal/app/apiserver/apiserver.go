package apiserver

import (
	"net/http"

	"github.com/Gugush284/Go-server.git/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// Set log level ...
func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// Configuration of router ...
func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

// New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start server ...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
