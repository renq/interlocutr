package comments

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateCommentRequest struct {
	Author string `json:"author"`
	Text string `json:"text"`
}

type CommentsResponse struct {
	Author string `json:"author"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

var storage []Comment

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
func GetComments(c echo.Context) error {
	result := make([]CommentsResponse, len(storage))
	
	for i, comment := range storage {
		result[i] = CommentsResponse{
			Author: comment.Author,
			Text: comment.Text,
			CreatedAt: comment.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, result)
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
func CreateComment(c echo.Context) error {
	comment := new(CreateCommentRequest)
	if err := c.Bind(comment); err != nil {
		return err
	}

	storage = append(storage, Comment{
		Author: comment.Author,
		Text: comment.Text,
		CreatedAt: time.Date(2026, 1, 6, 01, 12, 12, 0, time.UTC),
	})

	return c.JSON(http.StatusCreated, nil)
}

type Comment struct {
	Author string
	Text string
	CreatedAt time.Time
}
