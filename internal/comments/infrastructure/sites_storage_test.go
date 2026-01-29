package infrastructure_test

import (
	"fmt"
	"sync"
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

		err := storage.CreateSite(site)
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

	t.Run("concurrent create", func(t *testing.T) { createConcurrently(t, s, ids) })
	t.Run("concurrent get", func(t *testing.T) { getConcurrently(t, s, ids) })

	// Final consistency check
	for _, id := range ids {
		if _, err := s.GetSite(id); err != nil {
			t.Fatalf("missing site %s: %v", id, err)
		}
	}
}

func makeIDs(n int) []string {
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		ids[i] = fmt.Sprintf("id-%d", i)
	}
	return ids
}

func createConcurrently(t *testing.T, s app.SitesStorage, ids []string) {
	t.Helper()
	var wg sync.WaitGroup
	errCh := make(chan error, len(ids))

	for _, id := range ids {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.CreateSite(app.Site{ID: id}); err != nil {
				errCh <- err
			}
		}()
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Fatalf("create failed: %v", err)
	}
}

func getConcurrently(t *testing.T, s app.SitesStorage, ids []string) {
	t.Helper()
	var wg sync.WaitGroup
	errCh := make(chan error, len(ids))

	for _, id := range ids {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := s.GetSite(id); err != nil {
				errCh <- err
			}
		}()
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Fatalf("get failed: %v", err)
	}
}
