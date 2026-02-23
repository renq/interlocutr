package app

import (
	"time"

	"github.com/google/uuid"
	"github.com/renq/interlocutr/internal/infrastructure/clock"
	"github.com/renq/interlocutr/internal/infrastructure/uid"
)

type App struct {
	Clock           *clock.Clock
	CommentsStorage CommentsStorage
	SitesStorage    SitesStorage
	IDGenerator     uid.IDGenerator
}

func (a *App) FreezeTime(time time.Time) {
	a.Clock.FreezeTime(time)
}

func (a *App) GetNextIDValues(n int) []uuid.UUID {
	return a.IDGenerator.GetNextValues(n)
}

func NewApp(commentsStorage CommentsStorage, sitesStorage SitesStorage) *App {
	return &App{
		Clock:           clock.NewClock(),
		CommentsStorage: commentsStorage,
		SitesStorage:    sitesStorage,
	}
}
