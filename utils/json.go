package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Success bool        `json:"success,omitempty"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func NewJSONResponse(w http.ResponseWriter, params JSONResponse, headers ...http.Header) {
	respJSON, err := json.Marshal(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header().Set(k, v[0])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(params.Status)
	w.Write(respJSON)
}

func ParseJSONBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields in the JSON body

	if err := decoder.Decode(&data); err != nil {
		return err
	}

	err := decoder.Decode(&struct{}{}) // Check for trailing data
	if err != io.EOF {
		return errors.New("body must only contain a single JSON object")
	}

	return nil
}

func SendJSONError(w http.ResponseWriter, msg string, err error) JSONResponse {
	statusCode := http.StatusBadRequest

	return JSONResponse{
		Success: false,
		Status:  statusCode,
		Message: msg,
		Error:   err,
	}
}
