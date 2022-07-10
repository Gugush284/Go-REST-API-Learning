package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
	"github.com/Gugush284/Go-server.git/internal/app/store/teststore"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUserCreate(t *testing.T) {
	s := newServer(teststore.New())
	s.logger.SetLevel(logrus.ErrorLevel)

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
	u := model_user.TestUser(t)

	store := teststore.New()
	u, err := store.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	s := newServer(store)
	s.logger.SetLevel(logrus.ErrorLevel)

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
