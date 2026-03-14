package sqlite3_test

import (
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/infrastructure/sqlite3"
)

func createSitesStorage(t *testing.T) app.SitesStorage {
	db := connectToTestDB()
	// ctx := context.Background()

	// ctx := context.Background()
	// tx := db.MustBeginTx(ctx, nil)
	t.Cleanup(func() {
		// tx.Rollback()
		db.MustExec("DELETE FROM sites")
	})
	// << temp

	return sqlite3.NewSqliteSitesStorage(db)
}

func createCommentsStorage(t *testing.T) app.CommentsStorage {
	db := connectToTestDB()
	// ctx := context.Background()

	// ctx := context.Background()
	// tx := db.MustBeginTx(ctx, nil)
	t.Cleanup(func() {
		// tx.Rollback()
		db.MustExec("DELETE FROM comments")
	})
	// << temp

	return sqlite3.NewSqliteCommentsStorage(db)
}

func connectToTestDB() *sqlx.DB {
	appDB := os.Getenv("APP_DB")

	dsn := strings.TrimPrefix(appDB, "sqlite3://")
	db := sqlx.MustConnect("sqlite3", dsn)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(10)

	return db
}
