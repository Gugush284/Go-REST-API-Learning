package apiserver

import (
	"context"
	"encoding/json"
	"net/http"

	Constants "github.com/Gugush284/Go-server.git/internal/apiserver"
	model_user "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
)

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
			s.Err(w, r, http.StatusUnauthorized, Constants.ErrIncorrectLoginOrPassword)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		session, err := s.sessionStore.Get(r, Constants.SessionName)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Info("Request rejected as ", err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error("Request rejected as ", err)
			return
		}

		s.Logger.Info("Create an active session for ", u.ID)
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, Constants.SessionName)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.Err(w, r, http.StatusUnauthorized, Constants.ErrNotAuthenticated)
			s.Logger.Error(Constants.ErrNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.Err(w, r, http.StatusUnauthorized, Constants.ErrNotAuthenticated)
			s.Logger.Error(Constants.ErrNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), Constants.CtxKeyUser, u)))
	})
}
