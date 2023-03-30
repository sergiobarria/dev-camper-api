package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sergiobarria/dev-camper-api/api/handlers"
	"github.com/sergiobarria/dev-camper-api/utils"
	"github.com/spf13/viper"
)

func RegisterRoutes() *chi.Mux {
	debug := viper.GetBool("DEBUG")
	r := chi.NewRouter()

	// ====== REGISTER MIDDLEWARES ======
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	if debug {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer) // Recover must go after Logger

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// ====== REGISTER ROUTES ======
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		utils.NewJSONResponse(w, map[string]interface{}{
			// "success": true,
			"status":  http.StatusOK,
			"message": "Server is running",
		})
	})

	// ====== Bootcamps Routes ======
	r.Get("/bootcamps", handlers.GetBootcamps)
	r.Post("/bootcamps", handlers.CreateBootcamp)
	r.Get("/bootcamps/{id}", handlers.GetBootcamp)
	r.Patch("/bootcamps/{id}", handlers.UpdateBootcamp)
	r.Delete("/bootcamps/{id}", handlers.DeleteBootcamp)

	return r
}
