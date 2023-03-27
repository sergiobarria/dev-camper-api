package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sergiobarria/dev-camper-api/utils"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		utils.SendJSONResponse(w, http.StatusOK, utils.JSONResponse{
			Status:  "success",
			Message: "Server is running",
			Data:    nil,
		})
	})

	// Add other routes here 👇🏼

	return r
}
