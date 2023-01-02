package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @desc: Get all available bootcamps from the database
// @route: GET /api/v1/bootcamps
// @access: Public
func GetBootcamps(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"data": "Get all bootcamps",
	})
}

// @desc: Get a single bootcamp from the database
// @route: GET /api/v1/bootcamps/:id
// @access: Public
func GetBootcampById(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"data": "Get a single bootcamp",
		"id":   id,
	})
}

// @desc: Create a new bootcamp
// @route: POST /api/v1/bootcamps
// @access: Private
func CreateBootcamp(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"data": "Create a new bootcamp",
	})
}

// @desc: Update a bootcamp
// @route: PATCH /api/v1/bootcamps/:id
// @access: Private
func UpdateBootcamp(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"data": "Update a bootcamp",
		"id":   id,
	})
}

// @desc: Delete a bootcamp
// @route: DELETE /api/v1/bootcamps/:id
// @access: Private
func DeleteBootcamp(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusNoContent, map[string]string{
		"data": "Delete a bootcamp",
		"id":   id,
	})
}
