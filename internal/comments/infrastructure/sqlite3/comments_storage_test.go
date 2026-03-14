package sqlite3_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
)

func TestCommentsCanBeStoredAndRead(t *testing.T) {
	t.Parallel()
	t.Skip("not implemented yet")
	interfacestest.RunCommentsCanBeStoredAndReadTests(t, createCommentsStorage(t))
}

func TestBrokenStorage(t *testing.T) {
	t.Parallel()
	t.Skip("not implemented yet")
	interfacestest.RunBrokenStorageTests(t, createCommentsStorage(t))
}

func TestInMemoryCommentsStorage_ConcurrentCreateAndGet(t *testing.T) {
	t.Parallel()
	t.Skip("not implemented yet")
	interfacestest.RunCommentsStorageConcurrentTests(t, createCommentsStorage(t))
}
