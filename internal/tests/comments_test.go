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
	driver.FreezeTime(now)

	driver.LoginAsAdmin()

	// Arrange - add site
	createSiteResponse := driver.CreateSite(app.CreateSiteRequest{
		ID:      "test-site",
		Domains: []string{"interlocutr.lipek.net"},
	})
	assert.Equal(t, createSiteResponse.StatusCode, http.StatusCreated)

	// Act 1 - create comment
	createCommentResponse := driver.CreateComment(app.CreateCommentRequest{
		Site:     "test-site",
		Resource: "1",
		Author:   "Michał",
		Text:     "Jakiś tekst",
	})

	// Assert 1
	assert.Equal(t, http.StatusCreated, createCommentResponse.StatusCode)

	// Act 2
	getCommentsResponse := driver.GetComments("test-site", "1")

	// Assert 2
	assert.Equal(t, http.StatusOK, getCommentsResponse.StatusCode)

	assert.Equal(t, []app.GetCommentResponse{
		{
			Author:    "Michał",
			Text:      "Jakiś tekst",
			CreatedAt: now,
		}}, getCommentsResponse.Response)

	// req := httptest.NewRequest(http.MethodGet, "/api/test-site/1/comments", nil)
	// rec := httptest.NewRecorder()

	// expectedJson := `[{
	// 	"author": "Michał",
	// 	"text": "Jakiś tekst",
	// 	"created_at": "2026-01-06T01:12:12Z"
	// }]`

	// // Act
	// e.ServeHTTP(rec, req)

	// // Assert
	// assert.Equal(t, http.StatusOK, rec.Code)
	// assert.JSONEq(t, expectedJson, rec.Body.String())
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
