package sqlite3

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/renq/interlocutr/internal/comments/app"
)

type SqlxCommentsStorage struct {
	db *sqlx.DB
}

func NewSqliteCommentsStorage(db *sqlx.DB) app.CommentsStorage {
	return &SqlxCommentsStorage{
		db: db,
	}
}

func (s *SqlxCommentsStorage) CreateComment(ctx context.Context, comment app.Comment) error {
	return nil
}

func (s *SqlxCommentsStorage) GetComments(ctx context.Context, site, resource string) ([]app.Comment, error) {
	return []app.Comment{}, nil
}

func (s *SqlxCommentsStorage) Break() {

}
