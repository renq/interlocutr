package sqlite3

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/renq/interlocutr/internal/comments/app"
)

type SqlxCommentsStorage struct {
	db     *sqlx.DB
	broken bool
}

func NewSqliteCommentsStorage(db *sqlx.DB) app.CommentsStorage {
	return &SqlxCommentsStorage{
		db: db,
	}
}

func (s *SqlxCommentsStorage) CreateComment(ctx context.Context, comment app.Comment) error {
	if s.broken {
		return errors.New("storage is broken: can't store a new comment")
	}

	params := struct {
		ID        uuid.UUID `db:"id"`
		Site      string    `db:"site"`
		Resource  string    `db:"resource"`
		Author    string    `db:"author"`
		Text      string    `db:"text"`
		CreatedAt time.Time `db:"created_at"`
	}{
		ID:        comment.ID,
		Site:      comment.Site,
		Resource:  comment.Resource,
		Author:    comment.Author,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}

	query := `INSERT INTO comments (id, site, resource, author, text, created_at) 
	VALUES (:id, :site, :resource, :author, :text, :created_at)`
	_, err := s.db.NamedExecContext(ctx, query, params)

	// handle sqlite errors
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		// already exists
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
			return app.ErrorAlreadyExists
		}
		return errors.New("unexpected sqlite error: " + sqliteErr.Error())
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlxCommentsStorage) GetComments(ctx context.Context, site, resource string) ([]app.Comment, error) {
	comments := []app.Comment{}

	if s.broken {
		return comments, errors.New("storage is broken: can't read comments")
	}

	type row struct {
		ID        uuid.UUID `db:"id"`
		Site      string    `db:"site"`
		Resource  string    `db:"resource"`
		Author    string    `db:"author"`
		Text      string    `db:"text"`
		CreatedAt time.Time `db:"created_at"`
	}
	var rows []row

	err := s.db.SelectContext(ctx, &rows, "SELECT * FROM comments WHERE site = ? AND resource = ?", site, resource)

	if errors.Is(err, sql.ErrNoRows) {
		return comments, nil
	}

	if err != nil {
		return comments, err
	}

	comments = make([]app.Comment, len(rows))
	for i, r := range rows {
		comments[i] = app.Comment(r)
	}

	return comments, nil
}

func (s *SqlxCommentsStorage) Break() {
	s.broken = true
}
