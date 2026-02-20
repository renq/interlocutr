package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/renq/interlocutr/cmd"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/factory"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetComments(t *testing.T) {
	t.Parallel()

	// Arrange
	now, _ := time.Parse(time.RFC3339, "2026-01-06T01:12:12Z")

	driver := NewTestDriver(t)
	e := driver.E
	driver.FreezeTime(now)

	driver.LoginAsAdmin()

	// Arrange - add site
	response := driver.CreateSite(app.CreateSiteRequest{
		ID:      "test-site",
		Domains: []string{"interlocutr.lipek.net"},
	})
	assert.Equal(t, response.StatusCode, http.StatusCreated)

	// Arange 1
	createJson := `{
		"author": "Michał",
		"text": "Jakiś tekst"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/test-site/1/comments", strings.NewReader(createJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Act 1 - create comment
	e.ServeHTTP(rec, req)

	// Assert 1
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Act 2
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

func TestCommentCanBeAddedOnlyToValidSite(t *testing.T) {
	t.Parallel()

	app := factory.BuildApp()
	e := cmd.NewServer(app)

	createJson := `{
		"author": "Michał",
		"text": "Jakiś tekst"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/abcde/1/comments", strings.NewReader(createJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Act
	e.ServeHTTP(rec, req)

	// Assert 1
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
