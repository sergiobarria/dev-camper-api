package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ==================================================
// === Bootcamps Handlers - Path: /api/v1/products ===
// ==================================================

func GetBootcampsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "success",
		"data":   "Get all bootcamps",
	})
}

func GetBootcampHandler(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"status": "success",
		"data":   "Get single bootcamp",
		"id":     id,
	})
}

func CreateBootcampHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"status": "success",
		"data":   "Create new bootcamp",
	})
}

func UpdateBootcampHandler(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"status": "success",
		"data":   "Update bootcamp",
		"id":     id,
	})
}

func DeleteBootcampHandler(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusNoContent, map[string]string{
		"status": "success",
		"data":   "Delete bootcamp",
		"id":     id,
	})
}
