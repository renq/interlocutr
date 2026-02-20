package internal_tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/renq/interlocutr/cmd"
	"github.com/renq/interlocutr/internal/comments/factory"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuth(t *testing.T) {
	t.Parallel()

	t.Run("user can obtain JWT token if they use valid credentials", func(t *testing.T) {
		app := factory.BuildApp()
		e := cmd.NewServer(app)

		req := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader("username=admin&password=secret"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rec := httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEqual(t, "", bufferToJson(t, rec.Body)["token"])
	})

	t.Run("user can obtain JWT token by sending json playload", func(t *testing.T) {
		app := factory.BuildApp()
		e := cmd.NewServer(app)

		req := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(`{"username":"admin","password":"secret"}`))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEqual(t, "", bufferToJson(t, rec.Body)["token"])
	})

	t.Run("user needs to provide username and password, ex. as a form data", func(t *testing.T) {
		app := factory.BuildApp()
		e := cmd.NewServer(app)

		req := httptest.NewRequest(http.MethodPost, "/oauth/token", nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.PostForm = map[string][]string{
			"user": {"user"},
			"pass": {"password"},
		}
		rec := httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}
