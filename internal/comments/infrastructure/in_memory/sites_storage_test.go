package in_memory_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
	"github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
)

func TestNewInMemorySitesStorage(t *testing.T) {
	t.Parallel()
	interfacestest.RunSitesStorageTests(t, in_memory.NewInMemorySitesStorage())
}

func TestInMemorySitesStorage_ConcurrentCreateAndGet(t *testing.T) {
	t.Parallel()
	interfacestest.RunSitesStorageConcurrentTests(t, in_memory.NewInMemorySitesStorage())
}
