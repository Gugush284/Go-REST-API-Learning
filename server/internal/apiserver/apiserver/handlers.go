package apiserver

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	Constants "github.com/Gugush284/Go-server.git/internal/apiserver"
	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
	ModelUser "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
			return
		}

		u := &ModelUser.User{
			Login:             req.Login,
			DecryptedPassword: req.Password,
		}
		u, err := s.store.User().Create(u)
		if err != nil {
			s.Err(w, r, http.StatusUnprocessableEntity, err)
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

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Err(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByLogin(req.Login)
		if err != nil || !u.ComparePassword(req.Password) {
			s.Err(w, r, http.StatusUnauthorized, Constants.ErrIncorrectLoginOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, Constants.SessionName)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			return
		}

		s.Logger.Info("Create an active session for ", u.ID)
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) AuthenticateUser(next http.Handler) http.Handler {
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
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.Err(w, r, http.StatusUnauthorized, Constants.ErrNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), Constants.CtxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(Constants.CtxKeyUser).(*ModelUser.User))
	})
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuID := uuid.New().String()
		w.Header().Set("X-Request-ID", uuID)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), Constants.CtxKeyId, uuID)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.Logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(Constants.CtxKeyId),
		})
		logger.Infof("started %s, %s", r.Method, r.RequestURI)

		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof("completed with %d: %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start))
	})
}

func (s *server) UploadImage() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("image")
		if err != nil {
			s.Err(w, r, http.StatusUnprocessableEntity, err)
			s.Logger.Error(err)
			return
		}
		defer file.Close()

		path := strings.Join([]string{"assets", header.Filename}, "/")
		localFile, err := os.Create(path)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}
		defer localFile.Close()

		_, err = io.Copy(localFile, file)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		i := &ModelImage.Image{
			ImageName: header.Filename,
			Image:     path,
			Txt:       "-",
		}

		if err := s.store.Image().Upload(i); err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			if err = os.Remove(path); err != nil {
				s.Logger.Error(err)
			}
			return
		}

		s.respond(w, r, http.StatusCreated, i)
	})
}

func (s *server) Download() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := strings.ReplaceAll(r.URL.Path, "/download/", "")

		id, err := strconv.Atoi(key)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		i, err := s.store.Image().Download(id)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		fileBytes, err := ioutil.ReadFile(i.Image)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
	})
}
