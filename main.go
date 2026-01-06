package main

import (
	"interlocutr/comments"
	_ "interlocutr/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/:site/:resource/comments", comments.GetComments)
	e.POST("/:site/:resource/comments", comments.CreateComment)

	return e
}


// @title           Interlocutr API
// @version         1.0
// @contact.email  michal@lipek.net

// @license.name  MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	e := NewServer()

	e.Logger.Fatal(e.Start(":8080"))
}
