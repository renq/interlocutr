package app

import (
	"interlocutr/internal/infrastructure"
	"time"
)

type App struct {
	clock   *infrastructure.Clock
	storage Storage
}

func (a *App) FreezeTime(time time.Time) {
	a.clock.FreezeTime(time)
}

func NewApp() *App {
	return &App{
		clock:   infrastructure.NewClock(),
		storage: NewInMemoryStorage(),
	}
}
