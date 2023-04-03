package api

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sergiobarria/dev-camper-api/config"
	"github.com/sergiobarria/dev-camper-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	listenAddr *string
	debug      *bool
	client     *mongo.Client
	models     models.Models
}

func NewAPIServer(listendAddr *string, debug *bool, client *mongo.Client) *APIServer {
	return &APIServer{
		listenAddr: listendAddr,
		debug:      debug,
		client:     client,
		models:     models.NewModels(client.Database(config.EnvVars.MONGO_DB)),
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
	s.RegisterRoutes(app)

	mode := "development"
	if !*s.debug {
		mode = "production"
	}
	c := color.New(color.FgGreen, color.Bold, color.Underline)
	modeStr := c.Sprintf(mode)
	addrStr := c.Sprintf(":" + *s.listenAddr)

	fmt.Printf("⇨ 🚀 Http server running in %s mode on port %s \n", modeStr, addrStr)
	return app.Start(fmt.Sprintf(":%s", *s.listenAddr))
}
