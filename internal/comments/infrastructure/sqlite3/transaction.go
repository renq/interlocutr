package sqlite3

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type Modifier func(tx *sqlx.Tx) error
type Transaction func(db *sqlx.DB, modifier Modifier) error

func RunInTx(db *sqlx.DB, fn Modifier) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
