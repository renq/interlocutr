package app_test

import (
	"interlocutr/comments/app"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestSaveAndGetAll(t *testing.T) {
	t.Parallel()
	storage := app.InMemoryStorage{}

	comment1 := AComment()
	comment2 := AComment()
	assert.NoError(t, storage.CreateComment(comment1))
	assert.NoError(t, storage.CreateComment(comment2))

	comments, _ := storage.GetComments()
	assert.EqualValues(t, []app.Comment{comment1, comment2}, comments)
}

func TestBreakStorage(t *testing.T) {
	storage := app.InMemoryStorage{}
	storage.Break()

	t.Run("Broken storage returns error on creating a new comment", func(t *testing.T) {
		t.Parallel()

		assert.Error(t, storage.CreateComment(AComment()))
	})

	t.Run("Broken storage returns error on trying to read comments", func(t *testing.T) {
		t.Parallel()

		_, error := storage.GetComments()
		assert.Error(t, error)
	})
}

func AComment() app.Comment {
	return app.Comment{
		Author:    faker.FirstName(),
		Text:      faker.Sentence(),
		CreatedAt: time.Now().UTC(),
	}
}
