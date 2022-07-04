package apiserver

import (
	"io"
	"net/http"
)

// handler for '/hello'
func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
