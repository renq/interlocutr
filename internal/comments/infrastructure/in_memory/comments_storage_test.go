package in_memory_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
	infrastructure "github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
)

func TestCommentsCanBeStoredAndRead(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemoryCommentsStorage()

	interfacestest.RunCommentsCanBeStoredAndReadTests(t, storage)
}

func TestBrokenStorage(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemoryCommentsStorage()

	interfacestest.RunBrokenStorageTests(t, storage)
}

func TestInMemoryCommentsStorage_ConcurrentCreateAndGet(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemoryCommentsStorage()
	interfacestest.RunCommentsStorageConcurrentTests(t, storage)
}
