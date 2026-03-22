package sqlite3_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
)

func TestCommentsCanBeStoredAndRead(t *testing.T) {
	t.Parallel()
	interfacestest.RunCommentsCanBeStoredAndReadTests(t, createCommentsStorage(t))
}

func TestBrokenStorage(t *testing.T) {
	t.Parallel()
	interfacestest.RunBrokenStorageTests(t, createCommentsStorage(t))
}
