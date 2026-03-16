package sqlite3_test

import (
	"testing"

	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
)

func TestNewSqlStorage(t *testing.T) {
	t.Parallel()
	interfacestest.RunSitesStorageTests(t, createSitesStorage(t))
}
