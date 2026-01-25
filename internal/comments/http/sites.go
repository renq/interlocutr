package http

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	adminAuth "github.com/renq/interlocutr/internal/auth"
	"github.com/renq/interlocutr/internal/comments/app"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type SitesHandlers struct {
	app *app.App
}

func NewSitesHandlers(e *echo.Echo, app *app.App) SitesHandlers {
	h := SitesHandlers{
		app: app,
	}

	r := e.Group("/api/admin")
	config := echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(adminAuth.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))

	r.POST("/:site", h.CreateSite)

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
