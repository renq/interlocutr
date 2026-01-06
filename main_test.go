package main_test

import (
	"encoding/json"
	main "interlocutr"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestCreateAndGetComments(t *testing.T) {
	t.Parallel()

	// Arrange
	e := main.NewServer()

	req := httptest.NewRequest(http.MethodGet, "/test-site/1/comments", nil)
	rec := httptest.NewRecorder()

	expectedResponse := []main.CommentsResponse{
		{
			Author: "Michał",
			Text: "Jakiś tekst",
			CreatedAt: time.Date(2026, 1, 6, 01, 12, 12, 0, time.UTC),
		},
	}

	// Act
	e.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expectedResponse, getResponse[[]main.CommentsResponse](t, rec))
}


func getResponse[T any](t *testing.T, recorder *httptest.ResponseRecorder) T {
	var s *T
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &s))

	return *s
}
