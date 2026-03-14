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

func RunSitesStorageConcurrentTests(t *testing.T, storage app.SitesStorage) {
	// TODO Do we need these tests?
	ctx := context.Background()

	ids := makeIDs(20)

	t.Run("concurrent create", func(t *testing.T) {
		runConcurrently(t, ids, func(id string) error {
			_, err := storage.CreateSite(ctx, app.Site{ID: id})
			return err
		})
	})

	t.Run("concurrent get", func(t *testing.T) {
		runConcurrently(t, ids, func(id string) error {
			_, err := storage.GetSite(ctx, id)
			return err
		})
	})

	// Final consistency check
	for _, id := range ids {
		if _, err := storage.GetSite(ctx, id); err != nil {
			t.Fatalf("missing site %s: %v", id, err)
		}
	}
}
