package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(settings.EnvConfig.Port)))
}
