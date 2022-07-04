package apiserver

import "github.com/sirupsen/logrus"

// APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
	}
}

// Start server ...
func (s *APIserver) Start() error {
	return nil
}
