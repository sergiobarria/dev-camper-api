package main

import (
	"fmt"
	"net/http"
	"time"
)

func (a *application) listenAndServe() error {
	host := fmt.Sprintf("%s:%s", a.server.host, a.server.port)

	srv := http.Server{
		Handler:     a.router(),
		Addr:        host,
		ReadTimeout: 300 * time.Second,
	}

	a.infoLog.Printf("🚀 Server up and listening on %s\n", host)

	return srv.ListenAndServe()
}
