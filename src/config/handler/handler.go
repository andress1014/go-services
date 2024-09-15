package handler

import (
	"encoding/json"
	"net/http"
)

type ResponseApi struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
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
