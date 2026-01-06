package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/:site/:resource/comments", GetUser)
	e.Logger.Fatal(e.Start(":8080"))
}

type CommentsResponse struct {
	Author string `json:"author"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, []CommentsResponse{{
		Author: "Michał",
		Text: "Jakiś tekst",
		CreatedAt: time.Date(2026, 1, 6, 01, 12, 12, 0, time.UTC),
	}})
}
