package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sergiobarria/dev-camper-api/utils"
)

func GetBootcampsHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONResponse(w, http.StatusOK, utils.JSONResponse{
		Status: "success",
		Data:   "Get all bootcamps",
	})
}

func GetBootcampHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	utils.SendJSONResponse(w, http.StatusOK, utils.JSONResponse{
		Status: "success",
		Data:   "Get single bootcamp with ID: " + id,
	})
}

func CreateBootcampHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONResponse(w, http.StatusCreated, utils.JSONResponse{
		Status: "success",
		Data:   "Create new bootcamp",
	})
}

func UpdateBootcampHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	utils.SendJSONResponse(w, http.StatusOK, utils.JSONResponse{
		Status: "success",
		Data:   "Update bootcamp with ID: " + id,
	})
}

func DeleteBootcampHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	utils.SendJSONResponse(w, http.StatusOK, utils.JSONResponse{
		Status: "success",
		Data:   "Delete bootcamp with ID: " + id,
	})
}
