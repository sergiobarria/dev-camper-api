package router

import "github.com/labstack/echo/v4"

type ReviewsRouter struct{}

func NewReviewsRouter() ReviewsRouter {
	return ReviewsRouter{}
}

func (router *ReviewsRouter) ReviewsRoutes(e *echo.Group) {
	r := e.Group("/reviews")

	r.GET("", nil)
	r.GET("/:id", nil)
	r.POST("", nil)
	r.PUT("/:id", nil)
	r.DELETE("/:id", nil)
}
