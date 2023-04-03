package main

import (
	"context"
	"flag"
	"strconv"

	"github.com/sergiobarria/dev-camper-api/api"
	"github.com/sergiobarria/dev-camper-api/config"
)

func init() {
	config.LoadEnvVars() // Load environment variables
}

func main() {
	listenAddr := flag.String("port", strconv.Itoa(config.EnvVars.PORT), "Port to listen on")
	debug := flag.Bool("debug", config.EnvVars.DEBUG, "Enable debug mode")
	flag.Parse()

	// Connect to MongoDB
	client := config.NewMongoClient()
	defer client.Disconnect(context.TODO())

	// Create a new API server
	server := api.NewAPIServer(listenAddr, debug, client)

	// Run the server
	if err := server.Run(); err != nil {
		panic(err)
	}
}
