package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sergiobarria/dev-camper-api/models"
	"github.com/sergiobarria/dev-camper-api/repositories"
	"github.com/sergiobarria/dev-camper-api/utils"
)

type BootcampHandlers struct {
	repo repositories.BootcampRepository
}

func NewBootcampsHandlers(repo repositories.BootcampRepository) *BootcampHandlers {
	return &BootcampHandlers{repo}
}

func (h *BootcampHandlers) CreateBootcamp(w http.ResponseWriter, r *http.Request) {
	var bootcamp *models.Bootcamp

	err := utils.ParseJSONBody(w, r, &bootcamp)
	if err != nil {
		utils.NewJSONResponse(w, utils.JSONResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	bootcamp, _ = h.repo.InsertOne(bootcamp)

	utils.NewJSONResponse(w, utils.JSONResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Create new bootcamp",
		Data:    bootcamp,
	})
}

func (h *BootcampHandlers) GetBootcamps(w http.ResponseWriter, r *http.Request) {
	var bootcamps *[]models.Bootcamp

	bootcamps, err := h.repo.FindAll()
	if err != nil {
		utils.NewJSONResponse(w, utils.JSONResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Error getting bootcamps",
			Error:   err.Error(),
		})
		return
	}

	utils.NewJSONResponse(w, utils.JSONResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get all bootcamps",
		Data: map[string]interface{}{
			"count": len(*bootcamps),
			"data":  bootcamps,
		},
	})
}

func (h *BootcampHandlers) GetBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var bootcamp *models.Bootcamp

	bootcamp, err := h.repo.FindByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.NewJSONResponse(w, utils.JSONResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get bootcamp",
		Data:    bootcamp,
	})
}

func (h *BootcampHandlers) UpdateBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var bootcamp *models.Bootcamp

	err := utils.ParseJSONBody(w, r, &bootcamp)
	if err != nil {
		fmt.Println(err)
		return
	}

	bootcamp, err = h.repo.UpdateOne(id, *bootcamp)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.NewJSONResponse(w, utils.JSONResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Update bootcamp",
		Data:    bootcamp,
	})
}

func (h *BootcampHandlers) DeleteBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := h.repo.DeleteOne(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.NewJSONResponse(w, utils.JSONResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Delete bootcamp",
		Data:    map[string]string{"id": id},
	})
}
