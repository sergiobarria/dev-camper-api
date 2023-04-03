package main

import (
	"context"
	"flag"
	"strconv"

	"github.com/sergiobarria/dev-camper-api/api"
	"github.com/sergiobarria/dev-camper-api/initializers"
	"github.com/sergiobarria/dev-camper-api/store"
)

func main() {
	// Load environment variables from .env file
	cfg, err := initializers.LoadEnvVars()
	if err != nil {
		panic(err)
	}

	// Load & parse flags 👇🏼
	port := flag.String("port", ":"+strconv.Itoa(cfg.PORT), "port to listen on")
	debug := flag.Bool("debug", cfg.DEBUG, "debug mode")
	flag.Parse()

	// Connect to MongoDB 👇🏼
	client := initializers.NewMongoClient(cfg)
	defer client.Disconnect(context.Background())

	// Initialize models store 👇🏼
	store := store.New(client.Database(cfg.DB_NAME))

	// Create a new API server 👇🏼
	server := api.NewAPIServer(*port, debug, cfg, *store)

	// Start the server 👇🏼
	if err := server.Run(); err != nil {
		panic(err)
	}
}
