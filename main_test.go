package main_test

import (
	main "interlocutr"
	"interlocutr/comments/app"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetComments(t *testing.T) {
	t.Parallel()

	// Arrange
	now, _ := time.Parse(time.RFC3339, "2026-01-06T01:12:12Z")
	app := app.NewApp()
	app.FreezeTime(now)

	e := main.NewServer(app)

	createJson := `{
		"author": "Michał",
		"text": "Jakiś tekst"
	}`
	req := httptest.NewRequest(http.MethodPost, "/test-site/1/comments", strings.NewReader(createJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Act
	e.ServeHTTP(rec, req)

	// Assert 1
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Arrange 2
	req = httptest.NewRequest(http.MethodGet, "/test-site/1/comments", nil)
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

// func getResponse[T any](t *testing.T, recorder *httptest.ResponseRecorder) T {
// 	var s *T
// 	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &s))

// 	return *s
// }
