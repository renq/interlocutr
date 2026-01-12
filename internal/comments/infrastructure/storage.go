package infrastructure

import (
	"errors"

	"github.com/renq/interlocutr/internal/comments/app"
)

type bucket struct {
	Site     string
	Resource string
}

type storageStructure map[bucket][]app.Comment

type InMemoryStorage struct {
	storage storageStructure
	broken  bool
}

func NewInMemoryStorage() app.Storage {
	return &InMemoryStorage{
		storage: make(storageStructure),
	}
}

func (s *InMemoryStorage) CreateComment(comment app.Comment) error {
	if s.broken {
		return errors.New("storage is broken: can't store a new comment")
	}

	b := bucket{Site: comment.Site, Resource: comment.Resource}

	s.storage[b] = append(s.storage[b], comment)

	return nil
}

func (s *InMemoryStorage) GetComments(site, resource string) ([]app.Comment, error) {
	if s.broken {
		return []app.Comment{}, errors.New("storage is broken: can't read comments")
	}

	b := bucket{Site: site, Resource: resource}

	return s.storage[b], nil
}

func (s *InMemoryStorage) Break() {
	s.broken = true
}
