package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sergiobarria/dev-camper-api/models"
)

/**
* @desc: Bootcamps Routes - /api/v1/bootcamps
* @handle: CRUD Bootcamps
 */

func (s *APIServer) HandleCreateBootcamp(w http.ResponseWriter, r *http.Request) {
	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Status:  true,
		Message: "Create new bootcamp",
	})
}

func (s *APIServer) HandleGetBootcamps(w http.ResponseWriter, r *http.Request) {
	var bootcamps *[]models.Bootcamp

	bootcamps, err := s.bootcampRepo.FindAll()
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, JSONResponse{
			Status:  false,
			Message: "Error getting bootcamps",
			Error:   err.Error(),
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Status:  true,
		Message: "Show all bootcamps",
		Data: map[string]any{
			"count":     0,
			"bootcamps": bootcamps,
		},
	})
}

func (s *APIServer) HandleGetBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Status:  true,
		Message: "Show bootcamp",
		Data:    id,
	})
}

func (s *APIServer) HandleUpdateBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Status:  true,
		Message: "Update bootcamp",
		Data:    id,
	})
}

func (s *APIServer) HandleDeleteBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Status:  true,
		Message: "Delete bootcamp",
		Data:    id,
	})
}
