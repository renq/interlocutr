package main_test

import (
	"encoding/json"
	main "interlocutr"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func TestGetComments(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/test-site/1/comments", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := main.GetUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var got []main.CommentsResponse
	expected := []main.CommentsResponse{
		{
			Author: "Michał",
			Text: "Jakiś tekst",
			CreatedAt: time.Date(2026, 1, 6, 01, 12, 12, 0, time.UTC),
		},
	}

	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got))
	assert.Equal(t, expected, got)
}
