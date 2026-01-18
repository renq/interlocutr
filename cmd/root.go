/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/renq/interlocutr/docs"
	"github.com/renq/interlocutr/internal/comments/app"
	"github.com/renq/interlocutr/internal/comments/factory"
	"github.com/renq/interlocutr/internal/comments/http"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	http.NewCommentsHandlers(e, app)

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

	e.Logger.Fatal(e.Start(":" + port))
}
