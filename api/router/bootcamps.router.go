package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sergiobarria/dev-camper-api/api/handlers"
)

func RegiserBootcampsRoutes(e *echo.Group) {
	router := e.Group("/bootcamps")

	router.GET("", handlers.GetBootcampsHandler)
	router.GET("/:id", handlers.GetBootcampHandler)
	router.POST("", handlers.CreateBootcampHandler)
	router.PUT("/:id", handlers.UpdateBootcampHandler)
	router.DELETE("/:id", handlers.DeleteBootcampHandler)
}
