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
	IDs := driver.GetNextIDValues(1)

	driver.LoginAsAdmin()

	// Arrange - add site
	site := aSiteRequest()
	createSiteResponse := driver.CreateSite(site)
	assert.Equal(t, createSiteResponse.StatusCode, http.StatusCreated)

	// Act 1 - create comment
	newComment := aCommentRequest(site.ID, "1")
	createCommentResponse := driver.CreateComment(newComment)

	// Assert 1
	assert.Equal(t, http.StatusCreated, createCommentResponse.StatusCode)

	// Act 2
	getCommentsResponse := driver.GetComments(site.ID, "1")

	// Assert 2
	assert.Equal(t, http.StatusOK, getCommentsResponse.StatusCode)
	assert.Equal(t, []app.GetCommentResponse{
		{
			ID:        IDs[0],
			Author:    newComment.Author,
			Text:      newComment.Text,
			CreatedAt: now,
		}}, getCommentsResponse.Response)
}

func TestEachSiteHasTheirOwnComments(t *testing.T) {
	t.Parallel()

	// Arrange - add one comment to two different sites
	now, _ := time.Parse(time.RFC3339, "2026-01-06T01:12:12Z")

	driver := NewTestDriver(t)
	driver.FreezeTime(now)
	IDs := driver.GetNextIDValues(2)

	driver.LoginAsAdmin()

	site1 := aSiteRequest()
	site2 := aSiteRequest()

	createSite1Response := driver.CreateSite(site1)
	assert.Equal(t, createSite1Response.StatusCode, http.StatusCreated)
	createSite2Response := driver.CreateSite(site2)
	assert.Equal(t, createSite2Response.StatusCode, http.StatusCreated)

	site1Comment := aCommentRequest(site1.ID, "1")
	site2Comment := aCommentRequest(site2.ID, "1")

	createComment1Response := driver.CreateComment(site1Comment)
	createComment2Response := driver.CreateComment(site2Comment)

	assert.Equal(t, http.StatusCreated, createComment1Response.StatusCode)
	assert.Equal(t, http.StatusCreated, createComment2Response.StatusCode)

	// Act
	getSite1CommentsResponse := driver.GetComments(site1.ID, "1")
	getSite2CommentsResponse := driver.GetComments(site2.ID, "1")

	// Assert
	assert.Equal(t, http.StatusOK, getSite1CommentsResponse.StatusCode)
	assert.Equal(t, []app.GetCommentResponse{
		{
			ID:        IDs[0],
			Author:    site1Comment.Author,
			Text:      site1Comment.Text,
			CreatedAt: now,
		}}, getSite1CommentsResponse.Response)

	assert.Equal(t, http.StatusOK, getSite2CommentsResponse.StatusCode)
	assert.Equal(t, []app.GetCommentResponse{
		{
			ID:        IDs[1],
			Author:    site2Comment.Author,
			Text:      site2Comment.Text,
			CreatedAt: now,
		}}, getSite2CommentsResponse.Response)
}

func TestCommentCanBeAddedOnlyToValidSite(t *testing.T) {
	t.Parallel()

	driver := NewTestDriver(t)

	response := driver.CreateComment(aCommentRequest("non-existent-site", "1"))

	// Assert 1
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func aSiteRequest() app.CreateSiteRequest {
	return app.CreateSiteRequest{
		ID:      faker.Username(),
		Domains: []string{faker.DomainName()},
	}
}

func aCommentRequest(siteID string, resourceID string) app.CreateCommentRequest {
	return app.CreateCommentRequest{
		Site:     siteID,
		Resource: resourceID,
		Author:   faker.Name(),
		Text:     faker.Paragraph(),
	}
}
