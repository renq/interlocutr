package app

import (
	"time"

	"github.com/renq/interlocutr/internal/infrastructure/clock"
)

type App struct {
	Clock           *clock.Clock
	CommentsStorage CommentsStorage
	SitesStorage    SitesStorage
}

func (a *App) FreezeTime(time time.Time) {
	a.Clock.FreezeTime(time)
}

func NewApp(commentsStorage CommentsStorage, sitesStorage SitesStorage) *App {
	return &App{
		Clock:           clock.NewClock(),
		CommentsStorage: commentsStorage,
		SitesStorage:    sitesStorage,
	}
}
