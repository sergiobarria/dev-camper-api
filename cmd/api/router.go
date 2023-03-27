package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) router() http.Handler {
	r := chi.NewRouter()

	// ===== APPLY MIDDLEWARE 👇🏼 =====
	r.Use(middleware.RequestID) // Add a unique ID to each request
	r.Use(middleware.RealIP)    // Add the real IP address of the client to the request context
	r.Use(middleware.Recoverer) // Recover from panics without crashing the server

	if app.debug {
		r.Use(middleware.Logger) // Log the start and end of each request, only in debug mode (development)
	}

	// ===== APPLY ROUTES 👇🏼 =====
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(app.appName))
	})

	return r
}
