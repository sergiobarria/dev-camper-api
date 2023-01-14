package router

import "github.com/labstack/echo/v4"

type AuthRouter struct{}

func NewAuthRouter() AuthRouter {
	return AuthRouter{}
}

func (router *AuthRouter) AuthRoutes(e *echo.Group) {
	r := e.Group("/auth")

	r.POST("/register", nil)
	r.POST("/login", nil)
	r.GET("/me", nil)
	r.GET("/logout", nil)
	r.POST("/forgot-password", nil)
	r.PUT("/update-password", nil)
	r.PUT("/reset-password/:resetToken", nil)
	r.GET("/confirm-email", nil)
	r.PUT("/update-details", nil)
}
