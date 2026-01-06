package app

import "time"


type Comment struct {
	Author string
	Text string
	CreatedAt time.Time
}

type CreateCommentRequest struct {
	Author string `json:"author"`
	Text string `json:"text"`
}

type CommentsResponse struct {
	Author string `json:"author"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *App) GetComments() []CommentsResponse {
	result := make([]CommentsResponse, len(a.storage))
	
	for i, comment := range a.storage {
		result[i] = CommentsResponse{
			Author: comment.Author,
			Text: comment.Text,
			CreatedAt: comment.CreatedAt,
		}
	}

	return result
}

func (a *App) CreateComment(command CreateCommentRequest) error {
	a.storage = append(a.storage, Comment{
		Author: command.Author,
		Text: command.Text,
		CreatedAt: a.clock.Now(),
	})

	return nil
}
