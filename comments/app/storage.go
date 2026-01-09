package app

import "errors"

type bucket struct {
	Site     string
	Resource string
}

type storageStructure map[bucket][]Comment

type Storage interface {
	CreateComment(comment Comment) error
	GetComments(site, resource string) ([]Comment, error)
}

type InMemoryStorage struct {
	storage storageStructure
	broken  bool
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		storage: make(storageStructure),
	}
}

func (s *InMemoryStorage) CreateComment(comment Comment) error {
	if s.broken {
		return errors.New("Storage is broken. Can't store a new comment")
	}

	b := bucket{Site: comment.Site, Resource: comment.Resource}

	s.storage[b] = append(s.storage[b], comment)

	return nil
}

func (s *InMemoryStorage) GetComments(site, resource string) ([]Comment, error) {
	if s.broken {
		return []Comment{}, errors.New("Storage is broken. Can't read comments")
	}

	b := bucket{Site: site, Resource: resource}

	return s.storage[b], nil
}

func (s *InMemoryStorage) Break() {
	s.broken = true
}
