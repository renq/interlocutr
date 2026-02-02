package factory

import (
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/infrastructure"
)

func BuildApp() *app.App {
	return app.NewApp(
		infrastructure.NewInMemoryCommentsStorage(),
		infrastructure.NewInMemorySitesStorage(),
	)
}
