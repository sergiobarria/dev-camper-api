package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sergiobarria/dev-camper-api/api/router"
	"github.com/spf13/viper"
)

func StartServer() {
	e := echo.New()
	port := viper.GetInt("PORT")

	// Apply middlewares here 👇🏼
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Apply routes here 👇🏼
	router.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
