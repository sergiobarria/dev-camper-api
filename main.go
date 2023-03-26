package main

import (
	"github.com/sergiobarria/dev-camper-api/api"
	"github.com/sergiobarria/dev-camper-api/config"
)

func init() {
	// Load environment variables
	config.LoadEnvVars()
}

func main() {
	// Start server
	api.StartServer()
}
