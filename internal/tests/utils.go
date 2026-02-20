package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v5"

	"github.com/stretchr/testify/assert"
)

func bufferToJson(t *testing.T, body *bytes.Buffer) map[string]any {
	var responseBody map[string]any
	if e := json.Unmarshal(body.Bytes(), &responseBody); e != nil {
		assert.NoError(t, e, "response is not valid json: %s", body.String())
	}

	return responseBody
}

func bufferToStruct(t *testing.T, body *bytes.Buffer, out interface{}) {
	if e := json.Unmarshal(body.Bytes(), out); e != nil {
		assert.NoError(t, e)
	}
}

func getJWTToken(t *testing.T, e *echo.Echo) string {
	req := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(`{"username":"admin","password":"secret"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	token := bufferToJson(t, rec.Body)["token"].(string)
	assert.Equal(t, http.StatusOK, rec.Code)

	return token
}
