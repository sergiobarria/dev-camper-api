package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sergiobarria/dev-camper-api/api/v1/handlers"
)

func BootcampsRoutes() *chi.Mux {
	r := chi.NewRouter() // /api/v1/bootcamps

	r.Get("/", handlers.GetBootcampsHandler)
	r.Post("/", handlers.CreateBootcampHandler)
	r.Get("/{id}", handlers.GetBootcampHandler)
	r.Patch("/{id}", handlers.UpdateBootcampHandler)
	r.Delete("/{id}", handlers.DeleteBootcampHandler)

	return r
}
