package cmd

import (
	"log/slog"
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

func NewServer(app *app.App) *echo.Echo {
	e := echo.New()

	// Disabled because it does not work with echo v5 yet
	// e.GET("/swagger/*", echo.WrapHandler(echoSwagger.WrapHandler))

	// comments
	commentsHttp.NewCommentsHandlers(e, app)

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

	commentsHttp.NewSitesHandlers(admin, app)

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
