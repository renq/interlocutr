package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/renq/interlocutr/internal/comments/app"
)

type ApiResponse[T any] struct {
	StatusCode  int
	Response    T
	RawResponse map[string]any
}

type TestDriver struct {
	app      *app.App
	t        *testing.T
	e        *echo.Echo
	jwtToken string
}

func NewTestDriver(app *app.App, t *testing.T, e *echo.Echo) TestDriver {
	return TestDriver{
		app: app,
		t:   t,
		e:   e,
	}
}

func (d *TestDriver) LoginAsAdmin() {
	d.jwtToken = getJWTToken(d.t, d.e)
}

func (d *TestDriver) CreateSite(request app.CreateSiteRequest) ApiResponse[app.CreateSiteResponse] {
	req := httptest.NewRequest(http.MethodPost, "/api/admin/site", d.toBody(request))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.jwtToken)
	rec := httptest.NewRecorder()

	d.e.ServeHTTP(rec, req)

	var response app.CreateSiteResponse
	if rec.Code < 300 {
		bufferToStruct(d.t, rec.Body, &response)
	}
	return ApiResponse[app.CreateSiteResponse]{
		StatusCode:  rec.Code,
		Response:    response,
		RawResponse: bufferToJson(d.t, rec.Body),
	}
}

func (d *TestDriver) GetSite(siteID string) ApiResponse[app.GetSiteResponse] {
	req := httptest.NewRequest(http.MethodGet, "/api/admin/site/"+siteID, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.jwtToken)
	rec := httptest.NewRecorder()

	d.e.ServeHTTP(rec, req)

	var response app.GetSiteResponse
	if rec.Code < 300 {
		bufferToStruct(d.t, rec.Body, &response)
	}

	return ApiResponse[app.GetSiteResponse]{
		StatusCode:  rec.Code,
		Response:    response,
		RawResponse: bufferToJson(d.t, rec.Body),
	}
}

func (d *TestDriver) toBody(v any) *bytes.Reader {
	body, err := json.Marshal(v)
	if err != nil {
		d.t.Fatal(err)
	}

	return bytes.NewReader(body)
}
