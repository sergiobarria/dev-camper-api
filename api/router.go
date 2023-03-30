package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sergiobarria/dev-camper-api/api/handlers"
	"github.com/sergiobarria/dev-camper-api/repositories"
	"github.com/sergiobarria/dev-camper-api/utils"
	"github.com/spf13/viper"
)

func RegisterRoutes(repo repositories.BootcampRepository) *chi.Mux {
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
		utils.NewJSONResponse(w, utils.JSONResponse{
			Status:  http.StatusOK,
			Message: "Server is running",
		})
	})

	// ====== INSTANTIATE HANDLERS ======
	bootcampHandlers := handlers.NewBootcampsHandlers(repo)

	// ====== Bootcamps Routes ======
	r.Get("/bootcamps", bootcampHandlers.GetBootcamps)
	r.Post("/bootcamps", bootcampHandlers.CreateBootcamp)
	r.Get("/bootcamps/{id}", bootcampHandlers.GetBootcamp)
	r.Patch("/bootcamps/{id}", bootcampHandlers.UpdateBootcamp)
	r.Delete("/bootcamps/{id}", bootcampHandlers.DeleteBootcamp)

	return r
}
