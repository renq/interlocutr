package in_memory

import (
	"sync"

	"github.com/renq/interlocutr/internal/comments/app"
)

type sitesStorageStructure map[string]app.Site

type InMemorySitesStorege struct {
	mu      sync.RWMutex
	storage sitesStorageStructure
}

func NewInMemorySitesStorage() app.SitesStorage {
	return &InMemorySitesStorege{
		storage: make(sitesStorageStructure),
	}
}

func (s *InMemorySitesStorege) CreateSite(site app.Site) (app.Site, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.storage[site.ID]; ok {
		return app.Site{}, app.ErrorAlreadyExists
	}

	s.storage[site.ID] = site

	return site, nil
}

func (s *InMemorySitesStorege) GetSite(ID string) (app.Site, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	site, ok := s.storage[ID]
	if !ok {
		return app.Site{}, app.ErrorNotFound
	}

	return site, nil
}
