package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
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
	app      *app.App
	t        *testing.T
	e        *echo.Echo
	jwtToken string
}

func NewTestDriver(t *testing.T) TestDriver {
	app := factory.BuildApp()
	return TestDriver{
		app: app,
		t:   t,
		e:   cmd.NewServer(app),
	}
}

func (d *TestDriver) LoginAsAdmin() {
	d.jwtToken = getJWTToken(d.t, d.e)
}

func (d *TestDriver) FreezeTime(time time.Time) {
	d.app.FreezeTime(time)
}

func (d *TestDriver) GetNextIDValues(n int) []uuid.UUID {
	return d.app.GetNextIDValues(n)
}

// sites

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

// comments

func (d *TestDriver) CreateComment(request app.CreateCommentRequest) ApiResponse[app.CreateCommentResponse] {
	var payload struct {
		Author string `json:"author"`
		Text   string `json:"text"`
	}
	payload.Author = request.Author
	payload.Text = request.Text

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/%s/%s/comments", request.Site, request.Resource), d.toBody(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	d.e.ServeHTTP(rec, req)

	var response app.CreateCommentResponse
	if rec.Code < 300 {
		bufferToStruct(d.t, rec.Body, &response)
	}

	return ApiResponse[app.CreateCommentResponse]{
		StatusCode:  rec.Code,
		Response:    response,
		RawResponse: bufferToJson(d.t, rec.Body),
	}
}

func (d *TestDriver) GetComments(siteID string, resource string) ApiResponse[[]app.GetCommentResponse] {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/%s/%s/comments", siteID, resource), nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	d.e.ServeHTTP(rec, req)

	var response []app.GetCommentResponse
	if rec.Code < 300 {
		bufferToStruct(d.t, rec.Body, &response)
	}

	return ApiResponse[[]app.GetCommentResponse]{
		StatusCode: rec.Code,
		Response:   response,
		// RawResponse: bufferToJson(d.t, rec.Body),
	}
}

//

func (d *TestDriver) toBody(v any) *bytes.Reader {
	body, err := json.Marshal(v)
	if err != nil {
		d.t.Fatal(err)
	}

	return bytes.NewReader(body)
}
