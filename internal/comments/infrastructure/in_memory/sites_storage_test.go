package in_memory_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
	infrastructure "github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
)

func TestNewInMemorySitesStorage(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemorySitesStorage()

	interfacestest.RunSitesStorageTests(t, storage)
}

func TestInMemorySitesStorage_ConcurrentCreateAndGet(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemorySitesStorage()

	interfacestest.RunSitesStorageConcurrentTests(t, storage)
}
