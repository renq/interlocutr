package in_memory_test

import (
	"context"
	"testing"

	"github.com/renq/interlocutr/internal/comments/app"
	infrastructure "github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemorySitesStorage(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemorySitesStorage()
	ctx := context.Background()

	t.Run("returns error if site not found", func(t *testing.T) {
		_, err := storage.GetSite(ctx, "non-existing-site")
		assert.Equal(t, err, app.ErrorNotFound)
	})

	t.Run("returns site if it has been stored", func(t *testing.T) {
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

	t.Run("returns error if site with a given ID is already created", func(t *testing.T) {
		// Arrange
		siteID := "site2"
		site := app.Site{
			ID:      siteID,
			Domains: []string{"example.com", "example.org"},
		}

		_, err := storage.CreateSite(ctx, site)
		assert.NoError(t, err)

		// Act
		_, err = storage.CreateSite(ctx, site)

		// Assert
		assert.Equal(t, err, app.ErrorAlreadyExists)
	})
}

func TestInMemorySitesStorage_ConcurrentCreateAndGet(t *testing.T) {
	s := infrastructure.NewInMemorySitesStorage()
	ctx := context.Background()
	ids := makeIDs(20)

	t.Run("concurrent create", func(t *testing.T) {
		runConcurrently(t, ids, func(id string) error {
			_, err := s.CreateSite(ctx, app.Site{ID: id})
			return err
		})
	})

	t.Run("concurrent get", func(t *testing.T) {
		runConcurrently(t, ids, func(id string) error {
			_, err := s.GetSite(ctx, id)
			return err
		})
	})

	// Final consistency check
	for _, id := range ids {
		if _, err := s.GetSite(ctx, id); err != nil {
			t.Fatalf("missing site %s: %v", id, err)
		}
	}
}
