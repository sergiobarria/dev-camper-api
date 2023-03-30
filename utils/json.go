package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Success bool        `json:"success,omitempty"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func NewJSONResponse(w http.ResponseWriter, params map[string]interface{}, headers ...http.Header) {
	res := JSONResponse{
		Success: true,
		Status:  params["status"].(int),
		Data:    params["data"],
		Message: params["message"].(string),
		Error:   params["error"],
	}

	if v, ok := params["success"]; ok && v != nil {
		res.Success = v.(bool)
	}

	respJSON, err := json.Marshal(res)
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
	w.WriteHeader(res.Status)
	w.Write(respJSON)
}
