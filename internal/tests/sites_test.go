package tests

import (
	"net/http"
	"testing"

	"github.com/renq/interlocutr/cmd"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/factory"
	"github.com/stretchr/testify/assert"
)

func TestGetSites(t *testing.T) {
	t.Parallel()

	t.Run("site can't be created if request is not authorized", func(t *testing.T) {
		application := factory.BuildApp()
		e := cmd.NewServer(application)
		driver := NewTestDriver(application, t, e)

		createResponse := driver.CreateSite(app.CreateSiteRequest{
			ID:      "test-site",
			Domains: []string{"interlocutr.lipek.net"},
		})

		assert.Equal(t, http.StatusUnauthorized, createResponse.StatusCode)
	})

	t.Run("authorized user can add a get site info", func(t *testing.T) {
		application := factory.BuildApp()
		e := cmd.NewServer(application)

		// Arrange
		driver := NewTestDriver(application, t, e)
		driver.LoginAsAdmin()

		// Act
		createResponse := driver.CreateSite(app.CreateSiteRequest{
			ID:      "test-site",
			Domains: []string{"interlocutr.lipek.net"},
		})

		// Assert
		assert.Equal(t, http.StatusCreated, createResponse.StatusCode)
		assert.Equal(t, "test-site", createResponse.Response.ID)

		// Act
		getResponse := driver.GetSite("test-site")

		// Assert
		assert.Equal(t, http.StatusOK, getResponse.StatusCode)
		assert.Equal(t, app.GetSiteResponse{
			ID:      "test-site",
			Domains: []string{"interlocutr.lipek.net"},
		}, getResponse.Response)
	})
}
