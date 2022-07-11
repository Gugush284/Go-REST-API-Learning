package apiserver

import (
	"encoding/json"
	"net/http"

	globalErrors "github.com/Gugush284/Go-server.git/internal/apiserver"
	model_user "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
)

const sessionName = "activesession"

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("Request to create a user")

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Err(w, r, http.StatusBadRequest, err)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		u := &model_user.User{
			Login:             req.Login,
			DecryptedPassword: req.Password,
		}
		u, err := s.store.User().Create(u)
		if err != nil {
			s.Err(w, r, http.StatusUnprocessableEntity, err)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
		s.Logger.Info("Create user ", u.Login, " with id = ", u.ID)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("Request to create a session")

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Err(w, r, http.StatusBadRequest, err)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		u, err := s.store.User().FindByLogin(req.Login)
		if err != nil || !u.ComparePassword(req.Password) {
			s.Err(w, r, http.StatusUnauthorized, globalErrors.ErrIncorrectLoginOrPassword)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		s.Logger.Info("Create an active session for ", u.ID)
		s.respond(w, r, http.StatusOK, nil)
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
