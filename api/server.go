package api

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sergiobarria/dev-camper-api/initializers"
	"github.com/sergiobarria/dev-camper-api/store"
)

type APIServer struct {
	port   string
	debug  *bool
	config *initializers.Config
	store  store.StoreImpl
}

func NewAPIServer(port string, debug *bool, cfg *initializers.Config, store store.StoreImpl) *APIServer {
	return &APIServer{
		port:   port,
		debug:  debug,
		config: cfg,
		store:  store,
	}
}

func (s *APIServer) Run() error {
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	// ====== MIDDLEWARE ======
	if *s.debug {
		app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}))
	}
	app.Use(middleware.Recover())

	// ====== REGISTER ROUTES ======
	routerV1 := app.Group("/api/v1")

	routerV1.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status":  http.StatusOK,
			"message": "Server is running",
			"details": map[string]interface{}{
				"version": "1.0.0",
				"name":    "DevCamper API",
				"author":  "Sergio Barria",
				"license": "MIT",
			},
		})
	})

	routerV1.RouteNotFound("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, echo.Map{
			"success": false,
			"message": fmt.Sprintf("Route %s not found in this server", c.Request().URL.Path),
		})
	})

	routerV1.Any("*", func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, echo.Map{
			"success": false,
			"message": fmt.Sprintf("Method %s not allowed for route %s", c.Request().Method, c.Request().URL.Path),
		})
	})

	// Bootcamps routes
	routerV1.POST("/bootcamps", s.HandleCreateBootcamp)
	routerV1.GET("/bootcamps", s.HandleGetBootcamps)
	routerV1.GET("/bootcamps/:id", s.HandleGetBootcamp)
	routerV1.PATCH("/bootcamps/:id", s.HandleUpdateBootcamp)
	routerV1.DELETE("/bootcamps/:id", s.HandleDeleteBootcamp)

	// ====== START SERVER ======
	mode := s.config.GO_ENV
	c := color.New(color.FgGreen, color.Bold, color.Underline)
	modeStr := c.Sprintf(mode)
	addrStr := c.Sprintf(":" + s.port)

	fmt.Printf("⇨ 🚀 Http server running in %s mode on port %s \n", modeStr, addrStr)
	return app.Start(s.port)
}
