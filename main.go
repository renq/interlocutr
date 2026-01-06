package main

import (
	_ "interlocutr/docs"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/:site/:resource/comments", GetComments)

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

type CommentsResponse struct {
	Author string `json:"author"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// ShowAccount godoc
// @Summary      Get comments
// @Description  get comments for site and resource
// @Tags         comments
// @Accept       json
// @Produce      json
// @Success      200  {object}  []CommentsResponse
// Failure      400  {object}  httputil.HTTPError
// @Router       /{site}/{resource}/comments [get]
func GetComments(c echo.Context) error {
	return c.JSON(http.StatusOK, []CommentsResponse{{
		Author: "Michał",
		Text: "Jakiś tekst",
		CreatedAt: time.Date(2026, 1, 6, 01, 12, 12, 0, time.UTC),
	}})
}
