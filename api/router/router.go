package router

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	// API V1 group routes: /api/v1
	apiV1 := e.Group("/api/v1")

	/* Healthcheck route */
	apiV1.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "Ok!",
		})
	})

	// Add routes here 👇🏼
	RegiserBootcampsRoutes(apiV1) // /api/v1/bootcamps
}
