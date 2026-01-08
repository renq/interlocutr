package app

import "errors"

type Storage interface {
	CreateComment(comment Comment) error
	GetComments() ([]Comment, error)
}

type InMemoryStorage struct {
	storage []Comment
	broken  bool
}

func (s *InMemoryStorage) CreateComment(comment Comment) error {
	if s.broken {
		return errors.New("Storage is broken. Can't store a new comment")
	}

	s.storage = append(s.storage, comment)

	return nil
}

func (s *InMemoryStorage) GetComments() ([]Comment, error) {
	if s.broken {
		return []Comment{}, errors.New("Storage is broken. Can't read comments")
	}

	return s.storage, nil
}

func (s *InMemoryStorage) Break() {
	s.broken = true
}
