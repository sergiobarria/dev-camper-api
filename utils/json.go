package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Status  string      `json:"status"`
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)

	return nil
}
