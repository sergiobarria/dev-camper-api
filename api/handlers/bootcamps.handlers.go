package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sergiobarria/dev-camper-api/utils"
)

func CreateBootcamp(w http.ResponseWriter, r *http.Request) {
	utils.NewJSONResponse(w, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Create new bootcamp",
	})
}

func GetBootcamps(w http.ResponseWriter, r *http.Request) {

	utils.NewJSONResponse(w, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Show all bootcamps",
	})
}

func GetBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	utils.NewJSONResponse(w, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Show bootcamp",
		"data":    map[string]string{"id": id},
	})
}

func UpdateBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	utils.NewJSONResponse(w, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Update bootcamp",
		"data":    map[string]string{"id": id},
	})
}

func DeleteBootcamp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	utils.NewJSONResponse(w, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Delete bootcamp",
		"data":    map[string]string{"id": id},
	})
}
