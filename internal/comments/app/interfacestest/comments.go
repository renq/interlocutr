package interfacestest

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/stretchr/testify/assert"
)

func RunCommentsCanBeStoredAndReadTests(t *testing.T, storage app.CommentsStorage) {
	ctx := context.Background()

	t.Run("in a single page and post", func(t *testing.T) {
		comment1 := aComment("site1", "page1")
		comment2 := aComment("site1", "page1")
		assert.NoError(t, storage.CreateComment(ctx, comment1))
		assert.NoError(t, storage.CreateComment(ctx, comment2))

		comments, _ := storage.GetComments(ctx, "site1", "page1")
		assert.EqualValues(t, []app.Comment{comment1, comment2}, comments)
	})

	t.Run("can read comments only for a single site and post", func(t *testing.T) {
		comment1 := aComment("site2", "page2_1")
		comment2 := aComment("site3", "page3_1")
		assert.NoError(t, storage.CreateComment(ctx, comment1))
		assert.NoError(t, storage.CreateComment(ctx, comment2))

		comments, _ := storage.GetComments(ctx, "site2", "page2_1")
		assert.EqualValues(t, []app.Comment{comment1}, comments)
	})
}

func RunBrokenStorageTests(t *testing.T, storage app.CommentsStorage) {
	ctx := context.Background()
	storage.Break()

	t.Run("returns error on creating a new comment", func(t *testing.T) {
		t.Parallel()

		assert.Error(t, storage.CreateComment(ctx, aComment("any", "any")))
	})

	t.Run("returns error on trying to read comments", func(t *testing.T) {
		t.Parallel()

		_, error := storage.GetComments(ctx, "any", "any")
		assert.Error(t, error)
	})
}

func aComment(site, resource string) app.Comment {
	id, _ := uuid.NewV7()
	return app.Comment{
		ID:        id,
		Site:      site,
		Resource:  resource,
		Author:    faker.FirstName(),
		Text:      faker.Sentence(),
		CreatedAt: time.Now().UTC(),
	}
}
