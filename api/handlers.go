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
	var bootcamp models.Bootcamp

	err := ParseJSONBody(w, r, &bootcamp)
	if err != nil {
		SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// Create bootcamp
	err = s.bootcampRepo.InsertOne(&bootcamp)
	if err != nil {
		SendJSONError(w, http.StatusInternalServerError, err)
		return
	}

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Success: true,
		Data:    bootcamp,
	})
}

func (s *APIServer) HandleGetBootcamps(w http.ResponseWriter, r *http.Request) {
	var bootcamps *[]models.Bootcamp

	bootcamps, err := s.bootcampRepo.FindAll()
	if err != nil {
		SendJSONError(w, http.StatusInternalServerError, err)
		return
	}

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Success: true,
		Data: map[string]any{
			"count":     len(*bootcamps),
			"bootcamps": bootcamps,
		},
	})
}

func (s *APIServer) HandleGetBootcamp(w http.ResponseWriter, r *http.Request) {
	var bootcamp *models.Bootcamp
	id := chi.URLParam(r, "id")

	bootcamp, err := s.bootcampRepo.FindByID(id)
	if err != nil {
		SendJSONError(w, http.StatusInternalServerError, err)
		return
	}

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Success: true,
		Data:    bootcamp,
	})
}

func (s *APIServer) HandleUpdateBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var bootcamp models.Bootcamp

	err := ParseJSONBody(w, r, &bootcamp)
	if err != nil {
		SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// Update bootcamp
	err = s.bootcampRepo.UpdateOne(id, &bootcamp)
	if err != nil {
		SendJSONError(w, http.StatusInternalServerError, err)
		return
	}

	SendJSONResponse(w, http.StatusOK, JSONResponse{
		Success: true,
		Data:    bootcamp,
	})
}

func (s *APIServer) HandleDeleteBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := s.bootcampRepo.DeleteOne(id)
	if err != nil {
		SendJSONError(w, http.StatusInternalServerError, err)
		return
	}

	SendJSONResponse(w, http.StatusNoContent, JSONResponse{})
}
