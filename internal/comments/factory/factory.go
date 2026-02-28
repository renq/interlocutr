package factory

import (
	"github.com/renq/interlocutr/internal/comments/app"
	infrastructure "github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
)

func BuildApp() *app.App {
	return app.NewApp(
		infrastructure.NewInMemoryCommentsStorage(),
		infrastructure.NewInMemorySitesStorage(),
	)
}
