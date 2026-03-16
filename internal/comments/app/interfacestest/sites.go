package interfacestest

import (
	"context"
	"testing"

	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/stretchr/testify/assert"
)

func RunSitesStorageTests(t *testing.T, storage app.SitesStorage) {
	ctx := context.Background()

	t.Run("returns error if site not found", func(t *testing.T) {
		t.Parallel()
		_, err := storage.GetSite(ctx, "non-existing-site")
		assert.Equal(t, err, app.ErrorNotFound)
	})

	t.Run("returns site if it has been stored", func(t *testing.T) {
		t.Parallel()
		// Arrange
		siteID := "site1"
		site := app.Site{
			ID:      siteID,
			Domains: []string{"example.com", "example.org"},
		}

		_, err := storage.CreateSite(ctx, site)
		assert.NoError(t, err)

		// Act
		retrievedSite, err := storage.GetSite(ctx, siteID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, site, retrievedSite)
	})
}
