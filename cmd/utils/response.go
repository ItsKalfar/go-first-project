package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func SendResponse(w http.ResponseWriter, statusCode int, success bool, message string, data interface{}) {
    response := Response{
        Success: success,
        Message: message,
        Data:    data,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
