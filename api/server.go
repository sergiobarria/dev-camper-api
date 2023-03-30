package api

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sergiobarria/dev-camper-api/config"
	"github.com/sergiobarria/dev-camper-api/repositories"
	"github.com/spf13/viper"
)

type Server struct {
	listenAddr   string
	debug        bool
	infoLog      *log.Logger
	errLog       *log.Logger
	bootcampRepo repositories.BootcampRepository
}

func NewServer(listenAddr string) *Server {
	db := config.DB

	// Instantiate repositories
	bootcampRepo := repositories.NewBootcampRepo(db)

	return &Server{
		listenAddr:   listenAddr,
		debug:        viper.GetBool("DEBUG"),
		infoLog:      log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:       log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
		bootcampRepo: bootcampRepo,
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()

	// ====== Register routes ======
	r.Mount("/api/v1", RegisterRoutes(s.bootcampRepo))

	return http.ListenAndServe(":"+s.listenAddr, r)
}
