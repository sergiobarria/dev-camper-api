package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Status  bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type APIError struct {
	Error string `json:"error"`
}

func SendJSONResponse(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header().Set(k, v[0])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
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
