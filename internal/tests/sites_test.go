package internal_tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/renq/interlocutr/cmd"
	"github.com/renq/interlocutr/internal/comments/factory"
	"github.com/stretchr/testify/assert"
)

func TestGetSites(t *testing.T) {
	t.Parallel()

	t.Run("site can't be created if request is not authorized", func(t *testing.T) {
		app := factory.BuildApp()
		e := cmd.NewServer(app)

		req := httptest.NewRequest(http.MethodPost, "/api/admin/site", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("authorized user can add a get site info", func(t *testing.T) {
		app := factory.BuildApp()
		e := cmd.NewServer(app)

		// Step 1 - add site
		token := getJWTToken(t, e)

		req := httptest.NewRequest(http.MethodPost, "/api/admin/site", strings.NewReader(`{
			"ID": "test-site",
			"domains": ["interlocutr.lipek.net"]}
		`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response struct {
			ID string `json:"id"`
		}
		bufferToStruct(t, rec.Body, &response)
		assert.Equal(t, response.ID, "test-site")

		// Step 2 - get site info

		req = httptest.NewRequest(http.MethodGet, "/api/admin/site/test-site", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		rec = httptest.NewRecorder()

		// Act
		e.ServeHTTP(rec, req)

		// Assert
		assert.Equal(t, http.StatusOK, rec.Code)
		var getResponse struct {
			ID      string   `json:"id"`
			Domains []string `json:"domains"`
		}
		bufferToStruct(t, rec.Body, &getResponse)
		assert.Equal(t, "test-site", getResponse.ID)
		assert.Equal(t, []string{"interlocutr.lipek.net"}, getResponse.Domains)
	})
}
