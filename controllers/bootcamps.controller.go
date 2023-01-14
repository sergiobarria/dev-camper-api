package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sergiobarria/dev-camper-api/models"
	"github.com/sergiobarria/dev-camper-api/services"
)

type BootcampsController struct {
	service services.BootcampsService
}

func NewBootcampsController(service services.BootcampsService) BootcampsController {
	return BootcampsController{
		service: service,
	}
}

// @desc: Get all bootcamps from the database
// @route: GET /api/v1/bootcamps
// @access: Public
func (controller *BootcampsController) GetAllBootcamps(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid page",
		})
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid limit",
		})
	}

	sort := c.QueryParam("sort")
	selectFields := c.QueryParam("select")

	// Get all bootcamps
	bootcamps, err := controller.service.FindBootcamps(intPage, intLimit, sort, selectFields)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"results": len(bootcamps),
		"data":    bootcamps,
	})
}

// @desc: Get a single bootcamp from the database
// @route: GET /api/v1/bootcamps/:id
// @access: Public
func (controller *BootcampsController) GetBootcampById(c echo.Context) error {
	id := c.Param("id")

	bootcamp, err := controller.service.FindBootcampByID(id)
	if err != nil {
		if err == models.ErrNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"success": false,
				"error":   "Bootcamp not found",
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"success": false,
				"error":   "Server error",
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"data":    bootcamp,
	})
}

// @desc: Create a new bootcamp in the database
// @route: POST /api/v1/bootcamps
// @access: Private
func (controller *BootcampsController) CreateBootcamp(c echo.Context) error {
	var bootcamp *models.Bootcamp

	if err := c.Bind(&bootcamp); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid data",
		})
	}

	// Create a new bootcamp
	newBootcamp, err := controller.service.CreateBootcamp(bootcamp)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"error":   "Bootcamp already exists",
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"success": false,
				"error":   "Server error",
			})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"success": true,
		"data":    newBootcamp,
	})
}

// @desc: Update a bootcamp in the database
// @route: PUT /api/v1/bootcamps/:id
// @access: Private
func (controller *BootcampsController) UpdateBootcamp(c echo.Context) error {
	id := c.Param("id")

	var bootcamp *models.Bootcamp

	if err := c.Bind(&bootcamp); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"error":   "Invalid data",
		})
	}

	// Update a bootcamp
	updatedBootcamp, err := controller.service.UpdateBootcampByID(id, bootcamp)
	if err != nil {
		if err == models.ErrNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"success": false,
				"error":   "Bootcamp not found",
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"success": false,
				"error":   "Server error",
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"data":    updatedBootcamp,
	})
}

// @desc: Delete a bootcamp from the database
// @route: DELETE /api/v1/bootcamps/:id
// @access: Private
func (controller *BootcampsController) DeleteBootcamp(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeleteBootcampByID(id)

	if err != nil {
		if err == models.ErrNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"success": false,
				"error":   "Bootcamp not found",
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"success": false,
				"error":   "Server error",
			})
		}
	}

	return c.JSON(http.StatusNoContent, echo.Map{
		"success": true,
		"data":    nil,
	})
}
