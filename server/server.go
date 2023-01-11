package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func ListenAndServe() {
	e := echo.New() // create a new echo instance

	// apply middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[INFO]: method=${method}, uri=${uri}, status=${status}\n",
	}))

	// register route groups
	g := e.Group("/api/v1")

	// register handlers
	g.GET("/healthcheck", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// start the server
	port := viper.GetString("port") // get the port from the global config
	if port == "" {
		port = "1337"
	}

	e.Logger.Fatal(e.Start(":" + port)) // start the server
}
