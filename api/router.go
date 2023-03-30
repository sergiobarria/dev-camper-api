package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		utils.NewJSONResponse(w, map[string]interface{}{
			// "success": true,
			"status":  http.StatusOK,
			"message": "Server is running",
		})
	})

	// ====== Bootcamps ======
	r.Get("/bootcamps", handlers.GetBootcamps)
	r.Post("/bootcamps", handlers.CreateBootcamp)
	r.Get("/bootcamps/{id}", handlers.GetBootcamp)
	r.Patch("/bootcamps/{id}", handlers.UpdateBootcamp)
	r.Delete("/bootcamps/{id}", handlers.DeleteBootcamp)

	return r
}
