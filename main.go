package main

import (
	_ "github.com/renq/interlocutr/docs"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/factory"
	"github.com/renq/interlocutr/internal/comments/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewServer(app *app.App) *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	http.NewCommentsHandlers(e, app)

	return e
}

// @title           Interlocutr API
// @version         1.0
// @contact.email   michal@lipek.net
// @license.name    MIT

// @host            localhost:8080
// @BasePath        /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	e := NewServer(factory.BuildApp())

	e.Logger.Fatal(e.Start(":8080"))
}
