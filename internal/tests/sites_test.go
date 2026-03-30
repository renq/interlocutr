package tests

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/stretchr/testify/assert"
)

func TestGetSites(t *testing.T) {
	t.Parallel()

	t.Run("site can't be created if request is not authorized", func(t *testing.T) {
		driver := NewTestDriver(t)

		createResponse := driver.CreateSite(app.CreateSiteRequest{
			ID:      "anything",
			Domains: []string{"interlocutr.lipek.net"},
		})

		assert.Equal(t, http.StatusUnauthorized, createResponse.StatusCode)
	})

	t.Run("authorized user can add a get site info", func(t *testing.T) {
		// Arrange
		driver := NewTestDriver(t)
		driver.LoginAsAdmin()

		siteDomain := faker.DomainName()
		siteID := strings.ReplaceAll(siteDomain, ".", "-")

		// Act
		createResponse := driver.CreateSite(app.CreateSiteRequest{
			ID:      siteID,
			Domains: []string{siteDomain},
		})

		// Assert
		assert.Equal(t, http.StatusCreated, createResponse.StatusCode)
		assert.Equal(t, siteID, createResponse.Response.ID)

		// Act
		getResponse := driver.GetSite(siteID)

		// Assert
		assert.Equal(t, http.StatusOK, getResponse.StatusCode)
		assert.Equal(t, app.GetSiteResponse{
			ID:      siteID,
			Domains: []string{siteDomain},
		}, getResponse.Response)
	})
}
