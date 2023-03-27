package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Status  string      `json:"status"`
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}, headers ...http.Header) error {
	res, err := json.Marshal(data) // Convert data to JSON
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return err
	}

	// Set headers if any are passed in the arguments 👇🏼
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header().Set(k, v[0])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)

	return nil
}

func ReadJSONBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields in the request body
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	// Check there is only one JSON file in the request body
	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("request body must only contain one JSON object")
	}

	return nil
}

func SendJSONError(w http.ResponseWriter, err error, status ...int) error {
	var payload JSONResponse
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload.Status = "error"
	payload.Error = err.Error()

	return SendJSONResponse(w, statusCode, payload)
}
