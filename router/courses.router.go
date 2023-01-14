package router

import "github.com/labstack/echo/v4"

type CoursesRouter struct{}

func NewCoursesRouter() CoursesRouter {
	return CoursesRouter{}
}

func (router *CoursesRouter) ReviewsRoutes(e *echo.Group) {
	r := e.Group("/courses")

	r.GET("", nil)
	r.GET("/:id", nil)
	r.POST("", nil)
	r.PUT("/:id", nil)
	r.DELETE("/:id", nil)
}
