package infrastructure

import (
	"errors"

	"github.com/renq/interlocutr/internal/comments/app"
)

type bucket struct {
	Site     string
	Resource string
}

type commentsStorageStructure map[bucket][]app.Comment

type InMemoryCommentsStorage struct {
	storage commentsStorageStructure
	broken  bool
}

func NewInMemoryCommentsStorage() app.CommentsStorage {
	return &InMemoryCommentsStorage{
		storage: make(commentsStorageStructure),
	}
}

func (s *InMemoryCommentsStorage) CreateComment(comment app.Comment) error {
	if s.broken {
		return errors.New("storage is broken: can't store a new comment")
	}

	b := bucket{Site: comment.Site, Resource: comment.Resource}

	s.storage[b] = append(s.storage[b], comment)

	return nil
}

func (s *InMemoryCommentsStorage) GetComments(site, resource string) ([]app.Comment, error) {
	if s.broken {
		return []app.Comment{}, errors.New("storage is broken: can't read comments")
	}

	b := bucket{Site: site, Resource: resource}

	return s.storage[b], nil
}

func (s *InMemoryCommentsStorage) Break() {
	s.broken = true
}
