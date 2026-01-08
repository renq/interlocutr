package app

import "time"

type Comment struct {
	Author    string
	Text      string
	CreatedAt time.Time
}

type CreateCommentRequest struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type CommentsResponse struct {
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *App) GetComments() []CommentsResponse {
	comments, _ := a.storage.GetComments()
	result := make([]CommentsResponse, len(comments))

	for i, comment := range comments {
		result[i] = CommentsResponse(comment)
	}

	return result
}

func (a *App) CreateComment(command CreateCommentRequest) error {
	return a.storage.CreateComment(Comment{
		Author:    command.Author,
		Text:      command.Text,
		CreatedAt: a.clock.Now(),
	})
}
