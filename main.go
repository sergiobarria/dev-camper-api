package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sergiobarria/dev-camper-api/config"
	"github.com/sergiobarria/dev-camper-api/controllers"
	"github.com/sergiobarria/dev-camper-api/router"
	"github.com/sergiobarria/dev-camper-api/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *echo.Echo
	ctx         context.Context
	mongoClient *mongo.Client

	// Bootcamp Module (routes, controllers, services, models)
	bootcampsColl       *mongo.Collection
	bootcampsService    services.BootcampsService
	BootcampsController controllers.BootcampsController
	BootcampsRouter     router.BootcampsRouter

	// More global variables here 👇🏼
	// see: https://github.com/wpcodevo/golang-mongodb-api/blob/golang-mongodb-crud-api/cmd/server/main.go
)

func init() {
	cfg, err := config.LoadConfig(".") // Load global config
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to the database
	mongoConn := options.Client().ApplyURI(cfg.DBUri)
	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal("Could not connect to MongoDB", err)
	} else {
		log.Println("Connected to MongoDB")
	}

	// Ping the database to check if it is alive
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Could not ping MongoDB", err)
	}

	// Register Bootcamps Modules
	bootcampsColl = mongoClient.Database(cfg.DbName).Collection("bootcamps")
	bootcampsService = services.NewBootcampsService(bootcampsColl, ctx)
	BootcampsController = controllers.NewBootcampsController(bootcampsService)
	BootcampsRouter = router.NewBootcampsRouter(BootcampsController)

	// TODO: Register collections and services
	// see: https://github.com/wpcodevo/golang-mongodb-api/blob/golang-mongodb-crud-api/cmd/server/main.go

	// Start the server
	server = echo.New()
}

func startServer(cfg config.Config) {
	// Apply middlewares
	server.Use(middleware.Recover())
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[INFO]: method=${method}, uri=${uri}, status=${status}\n",
	}))

	// enable CORS
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// New router
	routerV1 := server.Group("/api/v1")

	// TODO: Register routes
	routerV1.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "OK",
		})
	})

	BootcampsRouter.BootcampsRoutes(routerV1) // Register bootcamps routes

	// Start the server
	server.Logger.Fatal(server.Start(":" + cfg.Port))
}

func main() {
	cfg, err := config.LoadConfig(".") // Load global config
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Close the connection to the database when the main function ends
	defer mongoClient.Disconnect(ctx)

	startServer(cfg)
}
