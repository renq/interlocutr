package app_test

import (
	"interlocutr/internal/comments/app"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCommentsCanBeStoredAndRead(t *testing.T) {
	t.Parallel()
	storage := app.NewInMemoryStorage()

	t.Run("in a single page and post", func(t *testing.T) {
		comment1 := AComment("site1", "page1")
		comment2 := AComment("site1", "page1")
		assert.NoError(t, storage.CreateComment(comment1))
		assert.NoError(t, storage.CreateComment(comment2))

		comments, _ := storage.GetComments("site1", "page1")
		assert.EqualValues(t, []app.Comment{comment1, comment2}, comments)
	})

	t.Run("can read comments only for a single site and post", func(t *testing.T) {
		comment1 := AComment("site2", "page2_1")
		comment2 := AComment("site3", "page3_1")
		assert.NoError(t, storage.CreateComment(comment1))
		assert.NoError(t, storage.CreateComment(comment2))

		comments, _ := storage.GetComments("site2", "page2_1")
		assert.EqualValues(t, []app.Comment{comment1}, comments)
	})
}

func TestBrokenStorage(t *testing.T) {
	storage := app.InMemoryStorage{}
	storage.Break()

	t.Run("returns error on creating a new comment", func(t *testing.T) {
		t.Parallel()

		assert.Error(t, storage.CreateComment(AComment("any", "any")))
	})

	t.Run("returns error on trying to read comments", func(t *testing.T) {
		t.Parallel()

		_, error := storage.GetComments("any", "any")
		assert.Error(t, error)
	})
}

func AComment(site, resource string) app.Comment {
	return app.Comment{
		Site:      site,
		Resource:  resource,
		Author:    faker.FirstName(),
		Text:      faker.Sentence(),
		CreatedAt: time.Now().UTC(),
	}
}
