package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/renq/interlocutr/internal/comments/app"
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

	// Arrange - new comment

	newComment := app.CreateCommentRequest{
		Site:     "test-site",
		Resource: "1",
		Author:   faker.Name(),
		Text:     faker.Paragraph(),
	}

	// Act 1 - create comment
	createCommentResponse := driver.CreateComment(newComment)

	// Assert 1
	assert.Equal(t, http.StatusCreated, createCommentResponse.StatusCode)

	// Act 2
	getCommentsResponse := driver.GetComments("test-site", "1")

	// Assert 2
	assert.Equal(t, http.StatusOK, getCommentsResponse.StatusCode)
	assert.Equal(t, []app.GetCommentResponse{
		{
			Author:    newComment.Author,
			Text:      newComment.Text,
			CreatedAt: now,
		}}, getCommentsResponse.Response)
}

func TestCommentCanBeAddedOnlyToValidSite(t *testing.T) {
	t.Parallel()

	driver := NewTestDriver(t)

	response := driver.CreateComment(app.CreateCommentRequest{
		Site:     "non-existent-site",
		Resource: "1",
		Author:   "any",
		Text:     "any",
	})

	// Assert 1
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
