package main

import (
	"github.com/sergiobarria/dev-camper-api/api"
	"github.com/sergiobarria/dev-camper-api/initializers"
)

func init() {
	initializers.LoadEnvVars()      // Load environment variables from .env file
	initializers.ConnectToMongoDB() // Connect to MongoDB
}

func main() {
	// Create Server
	server := api.NewServer(initializers.DB)

	// Start Server
	server.Run()
}
