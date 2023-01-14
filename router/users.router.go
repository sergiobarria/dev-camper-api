package router

import "github.com/labstack/echo/v4"

type UsersRouter struct{}

func NewUsersRouter() UsersRouter {
	return UsersRouter{}
}

func (router *UsersRouter) UsersRoutes(e *echo.Group) {
	r := e.Group("/users")

	r.GET("", nil)
	r.GET("/:id", nil)
	r.POST("", nil)
	r.PUT("/:id", nil)
	r.DELETE("/:id", nil)
}
