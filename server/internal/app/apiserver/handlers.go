package apiserver

import (
	"encoding/json"
	"net/http"

	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
)

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Err(w, r, http.StatusBadRequest, err)
			s.logger.Info(err)
			return
		}

		u := &model_user.User{
			Login:             req.Login,
			DecryptedPassword: req.Password,
		}
		u, err := s.store.User().Create(u)
		if err != nil {
			s.Err(w, r, http.StatusUnprocessableEntity, err)
			s.logger.Info(err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
		s.logger.Info("Create user ", u.Login, " with id = ", u.ID)
	}
}

func (s *server) Err(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
