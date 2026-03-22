package factory

import (
	"log/slog"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
	"github.com/renq/interlocutr/internal/comments/infrastructure/sqlite3"
)

func BuildApp(dbDSN string) *app.App {
	if dbDSN == "" {
		slog.Info("Starting app with in-memory database")
		return app.NewApp(
			in_memory.NewInMemoryCommentsStorage(),
			in_memory.NewInMemorySitesStorage(),
		)
	} else {
		slog.Info("Starting app with sql database")
		db := connectToDB(dbDSN)
		return app.NewApp(
			sqlite3.NewSqliteCommentsStorage(db),
			sqlite3.NewSqliteSitesStorage(db),
		)
	}
}

func connectToDB(dsn string) *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", strings.TrimPrefix(dsn, "sqlite3://"))
	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(10)

	return db
}
