package http

import (
	"interlocutr/comments/app"
	"net/http"

	"github.com/labstack/echo/v4"
)


type CommentsHandlers struct {
	app *app.App
}

func NewCommentsHandlers(e *echo.Echo, app *app.App) CommentsHandlers {
	h := CommentsHandlers{
		app: app,
	}

	e.GET("/:site/:resource/comments", h.GetComments)
	e.POST("/:site/:resource/comments", h.CreateComment)

	return h
}

// GetComments godoc
// @Summary      Get comments
// @Description  get comments for site and resource
// @Tags         comments
// @Produce      json
// @Param        site      path  string  true  "Site identifier"
// @Param        resource  path  string  true  "Resource identifier"
// @Success      200       {object}  []CommentsResponse
// Failure      400       {object}  echo.HTTPError
// @Router       /{site}/{resource}/comments [get]
func (h *CommentsHandlers) GetComments(c echo.Context) error {
	return c.JSON(http.StatusOK, h.app.GetComments())
}

// CreateComment godoc
// @Summary      Create comment
// @Description  Create comment for site and resource
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        site      path  string               true  "Site identifier"
// @Param        resource  path  string               true  "Resource identifier"
// @Param        comment   body  CreateCommentRequest true  "Comment to create"
// @Success      201       {object}  CommentsResponse
// Failure      400       {object}  echo.HTTPError
// @Router       /{site}/{resource}/comments [post]
func (h *CommentsHandlers) CreateComment(c echo.Context) error {
	comment := new(app.CreateCommentRequest)
	if err := c.Bind(comment); err != nil {
		return err
	}

	h.app.CreateComment(*comment)

	return c.JSON(http.StatusCreated, nil)
}
