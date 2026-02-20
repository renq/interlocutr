package cmd

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	token "github.com/renq/interlocutr/internal/auth"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/factory"
	commentsHttp "github.com/renq/interlocutr/internal/comments/http"
	"github.com/spf13/cobra"
)

var port string

var rootCmd = &cobra.Command{
	Use:   "interlocutr",
	Short: "Interlocurt is a simple comments service",
	Long: `To run interlocutr just run it without any parameters.
It will start the server on port 8080.
	
Run the app with --help to see all available options.`,
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&port, "port", "8080", "Port number for the server")
}

func NewServer(application *app.App) *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = func(c *echo.Context, err error) {
		if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
			if resp.Committed {
				return
			}
		}

		code := http.StatusInternalServerError
		message := "internal server error"

		switch {
		case errors.Is(err, app.ErrorNotFound):
			code = http.StatusNotFound
			message = err.Error()

		case errors.Is(err, app.ErrorAlreadyExists):
			code = http.StatusConflict
			message = err.Error()
		}

		var sc echo.HTTPStatusCoder
		if errors.As(err, &sc) { // find error in an error chain that implements HTTPStatusCoder
			if tmp := sc.StatusCode(); tmp != 0 {
				code = tmp
			}
		}

		_ = c.JSON(code, map[string]any{
			"error": message,
		})
	}

	// Disabled because it does not work with echo v5 yet
	// e.GET("/swagger/*", echo.WrapHandler(echoSwagger.WrapHandler))

	// public
	commentsHttp.NewCommentsHandlers(e, application)

	// admin
	token.NewTokenHandler(e)

	admin := e.Group("/api/admin")
	config := echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	admin.Use(echojwt.WithConfig(config))

	commentsHttp.NewSitesHandlers(admin, application)

	return e
}

// @title           Interlocutr API
// @version         1.0
// @contact.email   michal@lipek.net
// @license.name    MIT

// @host            localhost:8080
// @BasePath        /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func StartServer() {
	e := NewServer(factory.BuildApp())

	err := e.Start(":" + port)
	if err != nil {
		slog.Error("Error starting server", "error", err)
	}
}
