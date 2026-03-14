package sqlite3_test

import (
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/renq/interlocutr/internal/comments/app/interfacestest"
	"github.com/renq/interlocutr/internal/comments/infrastructure/sqlite3"
)

func TestNewSqlStorage(t *testing.T) {
	t.Parallel()

	db := ConnectToTestDB()
	// ctx := context.Background()

	// ctx := context.Background()
	// tx := db.MustBeginTx(ctx, nil)
	t.Cleanup(func() {
		// tx.Rollback()
		db.MustExec("DELETE FROM sites")
	})
	// << temp

	storage := sqlite3.NewSqliteSitesStorage(db)

	interfacestest.RunSitesStorageTests(t, storage)
}

func TestInSqlSitesStorage_ConcurrentCreateAndGet(t *testing.T) {
	db := ConnectToTestDB()
	// ctx := context.Background()

	// ctx := context.Background()
	// tx := db.MustBeginTx(ctx, nil)
	t.Cleanup(func() {
		// tx.Rollback()
		db.MustExec("DELETE FROM sites")
	})
	// << temp

	storage := sqlite3.NewSqliteSitesStorage(db)

	interfacestest.RunSitesStorageConcurrentTests(t, storage)
}

func ConnectToTestDB() *sqlx.DB {
	appDB := os.Getenv("APP_DB")

	dsn := strings.TrimPrefix(appDB, "sqlite3://")
	db := sqlx.MustConnect("sqlite3", dsn)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(10)

	return db
}
