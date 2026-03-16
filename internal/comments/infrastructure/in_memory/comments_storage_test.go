package in_memory_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
	"github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
)

func TestCommentsCanBeStoredAndRead(t *testing.T) {
	t.Parallel()
	interfacestest.RunCommentsCanBeStoredAndReadTests(t, in_memory.NewInMemoryCommentsStorage())
}

func TestBrokenStorage(t *testing.T) {
	t.Parallel()
	interfacestest.RunBrokenStorageTests(t, in_memory.NewInMemoryCommentsStorage())
}
