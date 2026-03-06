package factory

import (
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
)

func BuildApp() *app.App {
	return app.NewApp(
		in_memory.NewInMemoryCommentsStorage(),
		in_memory.NewInMemorySitesStorage(),
	)
}
