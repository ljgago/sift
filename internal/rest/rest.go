package rest

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// SuccessResponse returns the data.
func SuccessResponse(w http.ResponseWriter, code int, data any) {
	response := successResponse{
		Status: "success",
		Data:   data,
	}

	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)
}

type errorResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

// ErrorResponse returns a error message.
func ErrorResponse(w http.ResponseWriter, code int, message string) {
	response := errorResponse{
		Status:  "error",
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)
}
