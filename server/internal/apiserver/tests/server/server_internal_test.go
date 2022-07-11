package apiserver_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	Constants "github.com/Gugush284/Go-server.git/internal/apiserver"
	"github.com/Gugush284/Go-server.git/internal/apiserver/apiserver"
	"github.com/Gugush284/Go-server.git/internal/apiserver/store/teststore"
	ModelUserTest "github.com/Gugush284/Go-server.git/internal/apiserver/tests/model/users"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := ModelUserTest.TestUser(t)
	u, err := store.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	testcases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	s := apiserver.NewServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	fakeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodPost, "/", nil)
			cookieStr, _ := sc.Encode(Constants.SessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", Constants.SessionName, cookieStr))

			s.AuthenticateUser(fakeHandler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleUserCreate(t *testing.T) {
	s := apiserver.NewServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	s.Logger.SetLevel(logrus.ErrorLevel)

	testcases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"login": "us",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}

func TestServer_HandleSessionCreate(t *testing.T) {
	u := ModelUserTest.TestUser(t)

	store := teststore.New()
	u, err := store.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	s := apiserver.NewServer(store, sessions.NewCookieStore([]byte("se")))
	s.Logger.SetLevel(logrus.ErrorLevel)

	testcases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    u.Login,
				"password": u.DecryptedPassword,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "NotThisLogin",
			payload: map[string]string{
				"login":    "NotThisLogin",
				"password": u.DecryptedPassword,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "NotThisPassword",
			payload: map[string]string{
				"login":    u.Login,
				"password": "NotThisPassword",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Invalid payload",
			payload:      "Invalid payload",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}
