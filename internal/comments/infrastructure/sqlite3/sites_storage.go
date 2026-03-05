package sqlite3

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/renq/interlocutr/internal/comments/app"
)

type SqlxSitesStorege struct {
	db *sqlx.DB
}

func NewInSqlxSitesStorage(db *sqlx.DB) app.SitesStorage {
	return &SqlxSitesStorege{
		db: db,
	}
}

func (s *SqlxSitesStorege) CreateSite(site app.Site) (string, error) {
	domains, err := json.Marshal(site.Domains)
	if err != nil {
		return "", err // This should never happen, as domains is always a valid slice of strings.
	}

	params := struct {
		ID      string `db:"id"`
		Domains []byte `db:"domains"`
	}{
		ID:      site.ID,
		Domains: domains,
	}

	query := `INSERT INTO sites (id, domains) VALUES (:id, :domains)`
	// TODO Pass context to CreateSite
	_, err = s.db.NamedExecContext(context.Background(), query, params)

	// handle sqlite errors
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		// already exists
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
			return site.ID, app.ErrorAlreadyExists
		}
		return site.ID, errors.New("unexpected sqlite error: " + sqliteErr.Error())
	}

	if err != nil {
		return site.ID, err
	}

	return site.ID, nil
}

func (s *SqlxSitesStorege) GetSite(ID string) (app.Site, error) {
	site := app.Site{}

	var row struct {
		Id      string `db:"id"`
		Domains []byte `db:"domains"`
	}

	// TODO Pass context to GetSite
	err := s.db.GetContext(context.Background(), &row, "SELECT id, domains FROM sites WHERE id = ?", ID)

	if errors.Is(err, sql.ErrNoRows) {
		return site, app.ErrorNotFound
	}

	if err != nil {
		return site, err
	}

	site.ID = row.Id
	err = json.Unmarshal(row.Domains, &site.Domains)
	if err != nil {
		// should happen only if data in the database is invalid
		// log fatal error
		return site, err
	}

	return site, nil
}
