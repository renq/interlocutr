package app

import (
	"time"
)

type Comment struct {
	Site      string
	Resource  string
	Author    string
	Text      string
	CreatedAt time.Time
}

type CreateCommentRequest struct {
	Site     string `param:"site"`
	Resource string `param:"resource"`
	Author   string `json:"author"`
	Text     string `json:"text"`
}

type GetCommentResponse struct {
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *App) GetComments() []GetCommentResponse {
	comments, _ := a.CommentsStorage.GetComments("", "")
	result := make([]GetCommentResponse, len(comments))

	for i, comment := range comments {
		result[i] = GetCommentResponse{
			Author:    comment.Author,
			Text:      comment.Text,
			CreatedAt: comment.CreatedAt,
		}
	}

	return result
}

func (a *App) CreateComment(command CreateCommentRequest) error {
	_, err := a.SitesStorage.GetSite(command.Site)
	if err != nil {
		return err
	}

	return a.CommentsStorage.CreateComment(Comment{
		Author:    command.Author,
		Text:      command.Text,
		CreatedAt: a.Clock.Now(),
	})
}
