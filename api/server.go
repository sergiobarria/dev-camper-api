package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/sergiobarria/dev-camper-api/api/v1"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	appName  string
	debug    bool
	infoLog  *log.Logger
	errorLog *log.Logger
	router   *chi.Mux
	db       *mongo.Client

	// Add other dependencies here 👇🏼
}

func NewServer(db *mongo.Client) *Server {
	debug := viper.GetBool("DEBUG")
	appName := "DevCamper API V1"
	api := chi.NewRouter()

	// ===== Create Server =====
	server := &Server{
		appName:  appName,
		db:       db,
		router:   api,
		debug:    debug,
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile),

		// Add other dependencies here 👇🏼
	}

	// ===== APPLY MIDDLEWARES =====
	api.Use(middleware.RequestID) // Add a unique ID to each request
	api.Use(middleware.RealIP)    // Add the real IP address of the client to the request context
	api.Use(middleware.Recoverer) // Recover from panics without crashing the server

	if debug {
		api.Use(middleware.Logger) // Log the start and end of each request, only in debug mode (development)
	}

	// ===== APPLY ROUTES 👇🏼 =====
	api.Mount("/api/v1", v1.Router()) // Mount the v1 router

	return server
}

func (s *Server) Run() error {
	// ===== START SERVER =====
	host := fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT"))
	srv := http.Server{
		Handler:      s.router,
		Addr:         host,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a separate goroutine so that it doesn't block.
	go func() {
		s.infoLog.Printf("🚀 Server up and listening on %s\n", host)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.errorLog.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Wait for an interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	s.infoLog.Printf("🛑 %s is shutting down (signal: %s)", s.appName, sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.errorLog.Fatalf("Server forced to shutdown: %s", err)
	}

	return nil
}
