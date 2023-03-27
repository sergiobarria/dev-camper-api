package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sergiobarria/dev-camper-api/config"
	"github.com/spf13/viper"
)

type application struct {
	appName  string
	server   server
	debug    bool
	infoLog  *log.Logger
	errorLog *log.Logger
}

type server struct {
	host string
	port string
	url  string
}

func init() {
	config.LoadEnvVars()
}

func main() {
	port := viper.GetString("PORT")
	host := "localhost"

	server := server{
		host: host,
		port: port,
		url:  fmt.Sprintf("http://%s:%s", host, port),
	}

	app := application{
		appName:  "DevCamper API",
		server:   server,
		debug:    true,
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile),
	}

	if err := app.listenAndServe(); err != nil {
		app.errorLog.Fatal(err)
	}
}
