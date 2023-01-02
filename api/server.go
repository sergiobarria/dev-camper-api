package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sergiobarria/dev-camper-api/api/handlers"
	"github.com/sergiobarria/dev-camper-api/pkgs/settings"
)

// @desc: Init server function
func StartServer() {
	// Load env vars
	settings.LoadEnv()

	// Init Echo instance
	e := echo.New()

	// Apply middleware
	// ...

	// Apply routes
	// API V1
	api := e.Group("/api/v1")
	api.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// Bootcamps routes
	api.GET("/bootcamps", handlers.GetBootcamps)
	api.POST("/bootcamps", handlers.CreateBootcamp)
	api.GET("/bootcamps/:id", handlers.GetBootcampById)
	api.PATCH("/bootcamps/:id", handlers.UpdateBootcamp)
	api.DELETE("/bootcamps/:id", handlers.DeleteBootcamp)

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(settings.EnvConfig.Port)))
}
