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

	g.POST("/site", h.CreateSite)
	g.GET("/site/:id", h.GetSite)

	return h
}

// CreateSite godoc
// @Summary      Create site
// @Description  Create site
// @Tags         sites
// @Accept       json
// @Produce      json
// @Param        site      body      app.CreateSiteRequest  true  "Site"
// @Success      201       {object}  app.CreateSiteResponse
// @Failure      400       {object}  infrastructure.ErrorResponse
// @Failure      401       {object}  infrastructure.ErrorResponse
// @Router       /api/admin/site [post]
func (h *SitesHandlers) CreateSite(c *echo.Context) error {
	request := new(app.CreateSiteRequest)

	if err := c.Bind(request); err != nil {
		return err
	}

	response, err := h.app.CreateSite(*request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response)
}

// GetSite godoc
// @Summary      Get site
// @Description  Get site by id
// @Tags         sites
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Site ID"
// @Success      200  {object}  app.GetSiteResponse
// @Failure      400  {object}  infrastructure.ErrorResponse
// @Failure      401  {object}  infrastructure.ErrorResponse
// @Failure      404  {object}  infrastructure.ErrorResponse
// @Router       /api/admin/site/{id} [get]
func (h *SitesHandlers) GetSite(c *echo.Context) error {
	request := new(app.GetSiteRequest)

	if err := c.Bind(request); err != nil {
		return err
	}

	response, err := h.app.GetSite(*request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
