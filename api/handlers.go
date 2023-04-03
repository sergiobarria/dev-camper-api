package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) HandleCreateBootcamp(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Create bootcamp",
	})
}

func (s *APIServer) HandleGetBootcamps(c echo.Context) error {
	// s.store.Bootcamp.FindAll()
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Show all bootcamps",
	})
}

func (s *APIServer) HandleGetBootcamp(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Show bootcamp",
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
