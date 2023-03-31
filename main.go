package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
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
	debug := flag.String("debug", viper.GetString("DEBUG"), "Debug mode")
	flag.Parse()

	server := api.NewAPIServer(*listenAddr, debug)
	mode := "debug"
	if *debug == "false" {
		mode = "production"
	}
	c := color.New(color.FgGreen, color.Bold, color.Underline)
	modeStr := c.Sprintf(mode)
	addrStr := c.Sprintf(*listenAddr)

	fmt.Printf("🚀 Server running in %s mode on port %s \n", modeStr, addrStr)
	log.Fatal(server.Run())
}
