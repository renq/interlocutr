package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/renq/interlocutr/cmd"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/factory"
)

type ApiResponse[T any] struct {
	StatusCode  int
	Response    T
	RawResponse map[string]any
}

type TestDriver struct {
	App      *app.App // TODO make it privat
	t        *testing.T
	E        *echo.Echo // TODO make it privat
	jwtToken string
}

func NewTestDriver(t *testing.T) TestDriver {
	app := factory.BuildApp()
	return TestDriver{
		App: app,
		t:   t,
		E:   cmd.NewServer(app),
	}
}

func (d *TestDriver) LoginAsAdmin() {
	d.jwtToken = getJWTToken(d.t, d.E)
}

func (d *TestDriver) FreezeTime(time time.Time) {
	d.App.FreezeTime(time)
}

func (d *TestDriver) CreateSite(request app.CreateSiteRequest) ApiResponse[app.CreateSiteResponse] {
	req := httptest.NewRequest(http.MethodPost, "/api/admin/site", d.toBody(request))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.jwtToken)
	rec := httptest.NewRecorder()

	d.E.ServeHTTP(rec, req)

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

	d.E.ServeHTTP(rec, req)

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
