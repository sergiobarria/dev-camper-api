package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sergiobarria/dev-camper-api/api"
	"github.com/sergiobarria/dev-camper-api/config"
	"github.com/spf13/viper"
)

func init() {
	config.LoadEnvVars()      // Load environment variables from .env file
	config.ConnectToMongoDB() // Connect to MongoDB
}

func main() {
	listenAddr := flag.String("port", viper.GetString("PORT"), "Port to listen on")
	flag.Parse()

	// Create new server
	server := api.NewServer(*listenAddr)

	// Run server
	fmt.Println("🚀 Server running and listening on port:", *listenAddr)
	log.Fatal(server.Run())
}
