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

	// Bootcamps routes
	g.GET("/bootcamps", nil)
	g.GET("/bootcamps/:id", nil)
	g.POST("/bootcamps", nil)
	g.PUT("/bootcamps/:id", nil)
	g.DELETE("/bootcamps/:id", nil)

	// Courses routes
	g.GET("/courses", nil)
	g.GET("/courses/:id", nil)
	g.POST("/courses", nil)
	g.PUT("/courses/:id", nil)
	g.DELETE("/courses/:id", nil)

	// Auth routes
	g.POST("/auth/register", nil)
	g.POST("/auth/login", nil)
	g.GET("/auth/me", nil)
	g.GET("/auth/logout", nil)
	g.POST("/auth/forgot-password", nil)
	g.PUT("/auth/update-password", nil)
	g.PUT("/auth/reset-password/:resettoken", nil)
	g.GET("/auth/confirm-email", nil)
	g.PUT("/auth/update-details", nil)

	// Users routes
	g.GET("/users", nil)
	g.GET("/users/:id", nil)
	g.POST("/users", nil)
	g.PUT("/users/:id", nil)
	g.DELETE("/users/:id", nil)

	// Reviews routes
	g.GET("/reviews", nil)
	g.GET("/reviews/:id", nil)
	g.POST("/reviews", nil)
	g.PUT("/reviews/:id", nil)
	g.DELETE("/reviews/:id", nil)

	// start the server
	port := viper.GetString("port") // get the port from the global config
	if port == "" {
		port = "1337"
	}

	e.Logger.Fatal(e.Start(":" + port)) // start the server
}
