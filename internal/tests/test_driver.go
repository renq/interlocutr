package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/renq/interlocutr/cmd"
	token "github.com/renq/interlocutr/internal/auth"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/factory"
)

type Response struct {
	StatusCode int
	Body       *bytes.Buffer
}

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
	app := factory.BuildApp(os.Getenv("TEST_INTEGRATION_DB"))
	return TestDriver{
		app: app,
		t:   t,
		e:   cmd.NewServer(app),
	}
}

func (d *TestDriver) LoginAsAdmin() {
	response := d.GetJWTToken(token.LoginRequest{Username: "admin", Password: "secret"})
	d.jwtToken = response.Response.Token
}

func (d *TestDriver) FreezeTime(time time.Time) {
	d.app.FreezeTime(time)
}

func (d *TestDriver) GetNextIDValues(n int) []uuid.UUID {
	return d.app.GetNextIDValues(n)
}

// sites

func (d *TestDriver) CreateSite(request app.CreateSiteRequest) ApiResponse[app.CreateSiteResponse] {
	res := d.Request(
		http.MethodPost,
		"/api/admin/site",
		d.toBody(request),
		map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + d.jwtToken},
		},
	)

	var response app.CreateSiteResponse
	if res.StatusCode < 300 {
		bufferToStruct(d.t, res.Body, &response)
	}

	return ApiResponse[app.CreateSiteResponse]{
		StatusCode:  res.StatusCode,
		Response:    response,
		RawResponse: bufferToJson(d.t, res.Body),
	}
}

func (d *TestDriver) GetSite(siteID string) ApiResponse[app.GetSiteResponse] {
	res := d.Request(
		http.MethodGet,
		"/api/admin/site/"+siteID,
		nil,
		map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + d.jwtToken},
		},
	)

	var response app.GetSiteResponse
	if res.StatusCode < 300 {
		bufferToStruct(d.t, res.Body, &response)
	}

	return ApiResponse[app.GetSiteResponse]{
		StatusCode:  res.StatusCode,
		Response:    response,
		RawResponse: bufferToJson(d.t, res.Body),
	}
}

// comments

func (d *TestDriver) CreateComment(request app.CreateCommentRequest) ApiResponse[app.CreateCommentResponse] {
	res := d.Request(
		http.MethodPost,
		fmt.Sprintf("/api/%s/%s/comments", request.Site, request.Resource),
		d.toBody(struct {
			Author string `json:"author"`
			Text   string `json:"text"`
		}{Author: request.Author, Text: request.Text}),
		map[string][]string{"Content-Type": {"application/json"}},
	)

	var response app.CreateCommentResponse
	if res.StatusCode < 300 {
		bufferToStruct(d.t, res.Body, &response)
	}

	return ApiResponse[app.CreateCommentResponse]{
		StatusCode:  res.StatusCode,
		Response:    response,
		RawResponse: bufferToJson(d.t, res.Body),
	}
}

func (d *TestDriver) GetComments(siteID string, resource string) ApiResponse[[]app.GetCommentResponse] {
	res := d.Request(
		http.MethodGet,
		fmt.Sprintf("/api/%s/%s/comments", siteID, resource),
		nil,
		map[string][]string{"Content-Type": {"application/json"}},
	)

	var response []app.GetCommentResponse
	if res.StatusCode < 300 {
		bufferToStruct(d.t, res.Body, &response)
	}

	return ApiResponse[[]app.GetCommentResponse]{
		StatusCode: res.StatusCode,
		Response:   response,
		// RawResponse: bufferToJson(d.t, rec.Body),
	}
}

// jwt auth

func (d *TestDriver) GetJWTToken(request token.LoginRequest) ApiResponse[token.JwtResponse] {
	res := d.Request(
		http.MethodPost,
		"/oauth/token",
		strings.NewReader("username="+request.Username+"&password="+request.Password),
		map[string][]string{"Content-Type": {"application/x-www-form-urlencoded"}},
	)

	var response token.JwtResponse
	if res.StatusCode < 300 {
		bufferToStruct(d.t, res.Body, &response)
	}

	return ApiResponse[token.JwtResponse]{
		StatusCode:  res.StatusCode,
		Response:    response,
		RawResponse: bufferToJson(d.t, res.Body),
	}
}

//

func (d *TestDriver) Request(
	method string,
	url string,
	body io.Reader,
	headers map[string][]string,
) Response {
	req := httptest.NewRequest(method, url, body)
	for k, vals := range headers {
		for _, v := range vals {
			req.Header.Set(k, v)
		}
	}
	rec := httptest.NewRecorder()

	d.e.ServeHTTP(rec, req)

	return Response{
		StatusCode: rec.Code,
		Body:       rec.Body,
	}
}

func (d *TestDriver) toBody(v any) *bytes.Reader {
	body, err := json.Marshal(v)
	if err != nil {
		d.t.Fatal(err)
	}

	return bytes.NewReader(body)
}
