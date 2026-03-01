package in_memory_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/renq/interlocutr/internal/comments/app"
	infrastructure "github.com/renq/interlocutr/internal/comments/infrastructure/in_memory"
	"github.com/stretchr/testify/assert"
)

func TestCommentsCanBeStoredAndRead(t *testing.T) {
	t.Parallel()
	storage := infrastructure.NewInMemoryCommentsStorage()

	t.Run("in a single page and post", func(t *testing.T) {
		comment1 := aComment("site1", "page1")
		comment2 := aComment("site1", "page1")
		assert.NoError(t, storage.CreateComment(comment1))
		assert.NoError(t, storage.CreateComment(comment2))

		comments, _ := storage.GetComments("site1", "page1")
		assert.EqualValues(t, []app.Comment{comment1, comment2}, comments)
	})

	t.Run("can read comments only for a single site and post", func(t *testing.T) {
		comment1 := aComment("site2", "page2_1")
		comment2 := aComment("site3", "page3_1")
		assert.NoError(t, storage.CreateComment(comment1))
		assert.NoError(t, storage.CreateComment(comment2))

		comments, _ := storage.GetComments("site2", "page2_1")
		assert.EqualValues(t, []app.Comment{comment1}, comments)
	})
}

func TestBrokenStorage(t *testing.T) {
	storage := infrastructure.InMemoryCommentsStorage{}
	storage.Break()

	t.Run("returns error on creating a new comment", func(t *testing.T) {
		t.Parallel()

		assert.Error(t, storage.CreateComment(aComment("any", "any")))
	})

	t.Run("returns error on trying to read comments", func(t *testing.T) {
		t.Parallel()

		_, error := storage.GetComments("any", "any")
		assert.Error(t, error)
	})
}

func TestInMemoryCommentsStorage_ConcurrentCreateAndGet(t *testing.T) {
	s := infrastructure.NewInMemoryCommentsStorage()
	comments := makeComments(200, "site-conc", "page-conc")

	t.Run("concurrent create", func(t *testing.T) {
		runConcurrently(t, comments, func(c app.Comment) error {
			return s.CreateComment(c)
		})
	})

	t.Run("concurrent get", func(t *testing.T) {
		runConcurrently(t, comments, func(c app.Comment) error {
			got, err := s.GetComments(c.Site, c.Resource)
			if err != nil {
				return err
			}
			for _, g := range got {
				if g.Author == c.Author && g.Text == c.Text {
					return nil
				}
			}
			return fmt.Errorf("missing comment %s", c.Text)
		})
	})

	// Final deterministic check
	all, err := s.GetComments("site-conc", "page-conc")
	if err != nil {
		t.Fatalf("unexpected error reading final comments: %v", err)
	}
	if len(all) != len(comments) {
		t.Fatalf("expected %d comments, got %d", len(comments), len(all))
	}
}

func aComment(site, resource string) app.Comment {
	return app.Comment{
		Site:      site,
		Resource:  resource,
		Author:    faker.FirstName(),
		Text:      faker.Sentence(),
		CreatedAt: time.Now().UTC(),
	}
}

func makeComments(n int, site, resource string) []app.Comment {
	comments := make([]app.Comment, n)
	for i := 0; i < n; i++ {
		comments[i] = app.Comment{
			Site:      site,
			Resource:  resource,
			Author:    fmt.Sprintf("author-%d", i),
			Text:      fmt.Sprintf("text-%d", i),
			CreatedAt: time.Now().UTC(),
		}
	}
	return comments
}
