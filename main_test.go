package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/renq/interlocutr/cmd"
	"github.com/renq/interlocutr/internal/comments/factory"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetComments(t *testing.T) {
	t.Parallel()

	// Arrange
	now, _ := time.Parse(time.RFC3339, "2026-01-06T01:12:12Z")
	app := factory.BuildApp()
	app.FreezeTime(now)

	e := cmd.NewServer(app)

	createJson := `{
		"author": "Michał",
		"text": "Jakiś tekst"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/test-site/1/comments", strings.NewReader(createJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Act
	e.ServeHTTP(rec, req)

	// Assert 1
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Arrange 2
	req = httptest.NewRequest(http.MethodGet, "/api/test-site/1/comments", nil)
	rec = httptest.NewRecorder()

	expectedJson := `[{
		"author": "Michał",
		"text": "Jakiś tekst",
		"created_at": "2026-01-06T01:12:12Z"
	}]`

	// Act
	e.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expectedJson, rec.Body.String())
}

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

func TestGetSites(t *testing.T) {
	t.Parallel()
	// t.Skip("disabled")

	t.Run("site can't be created if request is not authorized", func(t *testing.T) {
		app := factory.BuildApp()
		e := cmd.NewServer(app)

		req := httptest.NewRequest(http.MethodPost, "/api/admin/test-site", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("authorized user can add a site", func(t *testing.T) {
		app := factory.BuildApp()
		e := cmd.NewServer(app)

		// //
		// req := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(`{"username":"admin","password":"secret"}`))
		// req.Header.Set("Content-Type", "application/json")
		// rec := httptest.NewRecorder()
		// e.ServeHTTP(rec, req)
		// token := bufferToJson(t, rec.Body)["token"].(string)
		// assert.Equal(t, http.StatusOK, rec.Code)
		// fmt.Println("Obtained token:", token)
		// //

		req := httptest.NewRequest(http.MethodPost, "/api/admin/test-site", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+getJWTToken(t, e))
		rec := httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusCreated, rec.Code)
	})
}

func bufferToJson(t *testing.T, body *bytes.Buffer) map[string]any {
	var responseBody map[string]any
	if e := json.Unmarshal(body.Bytes(), &responseBody); e != nil {
		assert.NoError(t, e)
	}

	return responseBody
}

func getJWTToken(t *testing.T, e *echo.Echo) string {
	req := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(`{"username":"admin","password":"secret"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	token := bufferToJson(t, rec.Body)["token"].(string)
	assert.Equal(t, http.StatusOK, rec.Code)

	return token
}
