package main

import (
	"github.com/sergiobarria/dev-camper-api/config"
	"github.com/sergiobarria/dev-camper-api/server"
)

func main() {
	err := config.LoadConfig() // Load global config
	if err != nil {
		panic(err) // If there is an error, panic
	}

	// Start the server
	server.ListenAndServe()
}
