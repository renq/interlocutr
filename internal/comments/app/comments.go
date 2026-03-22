package app

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID
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

type CreateCommentResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetCommentsRequest struct {
	Site     string `param:"site"`
	Resource string `param:"resource"`
}

type GetCommentResponse struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *App) GetComments(ctx context.Context, get GetCommentsRequest) []GetCommentResponse {
	comments, _ := a.CommentsStorage.GetComments(ctx, get.Site, get.Resource)
	result := make([]GetCommentResponse, len(comments))

	for i, comment := range comments {
		result[i] = GetCommentResponse{
			ID:        comment.ID,
			Author:    comment.Author,
			Text:      comment.Text,
			CreatedAt: comment.CreatedAt,
		}
	}

	return result
}

func (a *App) CreateComment(ctx context.Context, command CreateCommentRequest) (CreateCommentResponse, error) {
	_, err := a.SitesStorage.GetSite(ctx, command.Site)
	if err != nil {
		return CreateCommentResponse{}, err
	}

	ID, err := a.IDGenerator.Generate()
	if err != nil {
		return CreateCommentResponse{}, err
	}

	return CreateCommentResponse{ID: ID}, a.CommentsStorage.CreateComment(ctx, Comment{
		ID:        ID,
		Site:      command.Site,
		Resource:  command.Resource,
		Author:    command.Author,
		Text:      command.Text,
		CreatedAt: a.Clock.Now(),
	})
}
