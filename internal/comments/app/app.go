package app

import (
	"time"

	"github.com/renq/interlocutr/internal/infrastructure/clock"
)

type App struct {
	Clock           *clock.Clock
	CommentsStorage CommentsStorage
}

func (a *App) FreezeTime(time time.Time) {
	a.Clock.FreezeTime(time)
}

func NewApp(storage CommentsStorage) *App {
	return &App{
		Clock:           clock.NewClock(),
		CommentsStorage: storage,
	}
}
