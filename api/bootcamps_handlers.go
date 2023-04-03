package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sergiobarria/dev-camper-api/models"
)

func (s *APIServer) HandlerCreateBootcamp(c echo.Context) error {
	var b models.Bootcamp

	if err := c.Bind(&b); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err := s.models.Bootcamp.InsertOne(&b)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Create new bootcamp",
		"data":    b,
	})
}

func (s *APIServer) HandleFindBootcamps(c echo.Context) error {
	var b []models.Bootcamp

	b, err := s.models.Bootcamp.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Show all bootcamps",
		"data": map[string]interface{}{
			"count":     len(b),
			"bootcamps": b,
		},
	})
}

func (s *APIServer) HandleFindBootcamp(c echo.Context) error {
	id := c.Param("id")
	var b *models.Bootcamp

	b, err := s.models.Bootcamp.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Show single bootcamp",
		"data":    b,
	})
}

func (s *APIServer) HandleUpdateBootcamp(c echo.Context) error {
	id := c.Param("id")
	var b *models.Bootcamp

	if err := c.Bind(&b); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err := s.models.Bootcamp.UpdateOne(id, b)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Update bootcamp",
		"data":    b,
	})
}

func (s *APIServer) HandleDeleteBootcamp(c echo.Context) error {
	id := c.Param("id")

	err := s.models.Bootcamp.DeleteOne(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": "Bootcamp deleted",
		"data":    nil,
	})
}
