package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) HandlerCreateBootcamp(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Create new bootcamp",
		"data":    nil,
	})
}

func (s *APIServer) HandleFindBootcamps(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Show all bootcamps",
		"data":    nil,
	})
}

func (s *APIServer) HandleFindBootcamp(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Show single bootcamp",
		"data":    id,
	})
}

func (s *APIServer) HandleUpdateBootcamp(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Update bootcamp",
		"data":    id,
	})
}

func (s *APIServer) HandleDeleteBootcamp(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Delete bootcamp",
		"data":    id,
	})
}
