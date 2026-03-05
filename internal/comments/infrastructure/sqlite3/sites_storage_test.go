package sqlite3_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/infrastructure/sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestNewSqlStorage(t *testing.T) {
	t.Parallel()

	db := ConnectToTestDB()
	ctx := context.Background()

	// ctx := context.Background()
	// tx := db.MustBeginTx(ctx, nil)
	t.Cleanup(func() {
		// tx.Rollback()
		db.MustExec("DELETE FROM sites")
	})
	// << temp

	storage := sqlite3.NewInSqlxSitesStorage(db)

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

func TestInSqlSitesStorage_ConcurrentCreateAndGet(t *testing.T) {
	db := ConnectToTestDB()
	ctx := context.Background()

	// ctx := context.Background()
	// tx := db.MustBeginTx(ctx, nil)
	t.Cleanup(func() {
		// tx.Rollback()
		db.MustExec("DELETE FROM sites")
	})
	// << temp

	s := sqlite3.NewInSqlxSitesStorage(db)
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

func ConnectToTestDB() *sqlx.DB {
	appDB := os.Getenv("APP_DB")

	dsn := strings.TrimPrefix(appDB, "sqlite3://")
	db := sqlx.MustConnect("sqlite3", dsn)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(10)

	return db
}
