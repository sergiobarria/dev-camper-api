package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) RegisterRoutes(app *echo.Echo) {
	router := app.Group("/api/v1")

	router.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status":  http.StatusOK,
			"message": "Server is running",
			"details": map[string]interface{}{
				"version":    "1.0.0",
				"name":       "DevCamper API",
				"author":     "Sergio Barria",
				"license":    "MIT",
				"repository": "",
			},
		})
	})

	// ====== Bootcamps Routes ======
	router.GET("/bootcamps", s.HandleFindBootcamps)
	router.GET("/bootcamps/:id", s.HandleFindBootcamp)
	router.POST("/bootcamps", s.HandlerCreateBootcamp)
	router.PATCH("/bootcamps/:id", s.HandleUpdateBootcamp)
	router.DELETE("/bootcamps/:id", s.HandleDeleteBootcamp)

	router.RouteNotFound("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, echo.Map{
			"success": false,
			"message": fmt.Sprintf("Route %s not found in this server", c.Request().URL.Path),
		})
	})

	router.Any("*", func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, echo.Map{
			"success": false,
			"message": fmt.Sprintf("Method %s not allowed for route %s", c.Request().Method, c.Request().URL.Path),
		})
	})
}
