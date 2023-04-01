package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sergiobarria/dev-camper-api/repositories"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	listenAddr   string
	debug        *string
	bootcampRepo repositories.BootcampRepo
}

func NewAPIServer(listenAddr string, debug *string, client *mongo.Client) *APIServer {
	// ====== Register Repositories ======
	bootcampRepo := repositories.NewBootcampRepo(client)

	return &APIServer{
		listenAddr:   listenAddr,
		debug:        debug,
		bootcampRepo: bootcampRepo,
	}
}

func (s *APIServer) Run() error {
	debug := viper.GetBool("DEBUG")
	router := chi.NewRouter()

	// ====== APPLY MIDDLEWARES ======
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	if debug {
		router.Use(middleware.Logger) // logger must go before recoverer
	}
	router.Use(middleware.Recoverer)

	// TODO: Add globalErrorHandler middleware here 👇🏼

	router.Use(cors.Handler(cors.Options{
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
	router.Mount("/api/v1", s.RegisterRoutes())

	return http.ListenAndServe(":"+s.listenAddr, router)
}

func (s *APIServer) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	// Healtheck Route
	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		SendJSONResponse(w, http.StatusOK, JSONResponse{
			Success: true,
			Message: "DevCamper API v1.0.0 - Status: OK",
		})
	})

	// ====== Bootcamps Routes ======
	router.Get("/bootcamps", s.HandleGetBootcamps)
	router.Post("/bootcamps", s.HandleCreateBootcamp)
	router.Get("/bootcamps/{id}", s.HandleGetBootcamp)
	router.Patch("/bootcamps/{id}", s.HandleUpdateBootcamp)
	router.Delete("/bootcamps/{id}", s.HandleDeleteBootcamp)

	// ====== Other Routes ======
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		SendJSONResponse(w, http.StatusMethodNotAllowed, JSONResponse{
			Success: false,
			Message: fmt.Sprintf("Method %s not allowed", r.Method),
		})
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		SendJSONResponse(w, http.StatusNotFound, JSONResponse{
			Success: false,
			Message: "Route not found on this server",
		})
	})

	return router
}
