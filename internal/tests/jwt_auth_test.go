package tests

import (
	"net/http"
	"strings"
	"testing"

	token "github.com/renq/interlocutr/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuth(t *testing.T) {
	t.Parallel()

	t.Run("user can obtain JWT token if they use valid credentials", func(t *testing.T) {
		driver := NewTestDriver(t)

		tokenResponse := driver.GetJWTToken(token.LoginRequest{Username: "admin", Password: "secret"})

		assert.Equal(t, http.StatusOK, tokenResponse.StatusCode)
		assert.NotEqual(t, "", tokenResponse.Response.Token)
	})

	t.Run("user can obtain JWT token by sending json playload", func(t *testing.T) {
		driver := NewTestDriver(t)

		res := driver.Request(
			http.MethodPost,
			"/oauth/token",
			strings.NewReader(`{"username":"admin","password":"secret"}`),
			map[string][]string{"Content-Type": {"application/json"}},
		)

		// Assert
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.NotEqual(t, "", bufferToJson(t, res.Body)["token"])
	})

	t.Run("user needs to provide username and password, ex. as a form data", func(t *testing.T) {
		driver := NewTestDriver(t)
		res := driver.GetJWTToken(token.LoginRequest{Username: "user", Password: "password"})

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})
}
