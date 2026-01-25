package http

import (
	"net/http"

	"github.com/renq/interlocutr/internal/comments/app"

	"github.com/labstack/echo/v5"
)

type SitesHandlers struct {
	app *app.App
}

func NewSitesHandlers(g *echo.Group, app *app.App) SitesHandlers {
	h := SitesHandlers{
		app: app,
	}

	g.POST("/:site", h.CreateSite)

	return h
}

// CreateComment godoc
// @Summary      Create site
// @Description  Create site
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        site      path  string               true  "Site identifier"
// Success      201       {object}  app.CommentsResponse
// @Failure      400       {object}  infrastructure.ErrorResponse
// @Failure      401       {object}  infrastructure.ErrorResponse
// @Router       /api/admin/{site} [post]
func (h *SitesHandlers) CreateSite(c *echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
