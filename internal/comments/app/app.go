package app

import (
	"time"

	"github.com/renq/interlocutr/internal/infrastructure/clock"
)

type App struct {
	Clock   *clock.Clock
	Storage Storage
}

func (a *App) FreezeTime(time time.Time) {
	a.Clock.FreezeTime(time)
}

func NewApp(storage Storage) *App {
	return &App{
		Clock:   clock.NewClock(),
		Storage: storage,
	}
}
