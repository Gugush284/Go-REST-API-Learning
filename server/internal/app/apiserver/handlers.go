package apiserver

import (
	"io"
	"net/http"
)

func (s *server) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "User")
	}
}
