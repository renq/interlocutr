package cmd

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v5"
	_ "github.com/renq/interlocutr/docs"
	adminAuth "github.com/renq/interlocutr/internal/auth"
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
You can then access the API documentation under /swagger/index.html
	
Run the app with --help to see all available options.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

	commentsHttp.NewCommentsHandlers(e, app)
	adminAuth.NewAuthHandler(e)
	commentsHttp.NewSitesHandlers(e, app)

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
