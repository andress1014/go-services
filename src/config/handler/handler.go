package handler

import (
	"encoding/json"
	"net/http"
)

type ResponseApi struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func HandleResponse(w http.ResponseWriter, statusCode int, data any) {
	response := ResponseApi{
		Status: statusCode,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func HandleError(w http.ResponseWriter, statusCode int, message string) {
	response := ResponseApi{
		Status: statusCode,
		Error:  message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}
