package infrastructure_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemorySitesStorage(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemorySitesStorage()

	t.Run("returns error if site not found", func(t *testing.T) {
		_, err := storage.GetSite("non-existing-site")
		assert.Error(t, err, app.ErrorNotFound)
	})

	t.Run("returns site if it has been stored", func(t *testing.T) {
		// Arrange
		siteID := "site1"
		site := app.Site{
			ID:      siteID,
			Domains: []string{"example.com", "example.org"},
		}

		_, err := storage.CreateSite(site)
		assert.NoError(t, err)

		// Act
		retrievedSite, err := storage.GetSite(siteID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, site, retrievedSite)
	})
}

func TestInMemorySitesStorage_ConcurrentCreateAndGet(t *testing.T) {
	s := infrastructure.NewInMemorySitesStorage()
	ids := makeIDs(200)

	t.Run("concurrent create", func(t *testing.T) {
		runConcurrently(t, ids, func(id string) error {
			_, err := s.CreateSite(app.Site{ID: id})
			return err
		})
	})

	t.Run("concurrent get", func(t *testing.T) {
		runConcurrently(t, ids, func(id string) error {
			_, err := s.GetSite(id)
			return err
		})
	})

	// Final consistency check
	for _, id := range ids {
		if _, err := s.GetSite(id); err != nil {
			t.Fatalf("missing site %s: %v", id, err)
		}
	}
}
