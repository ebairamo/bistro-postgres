package handler

import (
	"bistro/models"
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, statusCode int, errorCode string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := models.ErrorResponse{
		Code:    errorCode,
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResponse)
}
