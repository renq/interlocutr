package main_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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

func TestGetSites(t *testing.T) {
	t.Parallel()
	t.Skip("Not implemented yet")

	t.Run("site can't be created if token is not present in the request", func(t *testing.T) {
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
}
