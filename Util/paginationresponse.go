package util

import (
	"encoding/json"
	"net/http"
)

type PaginatedResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int64       `json:"total"`
	Error   string      `json:"error,omitempty"`
}

// SendPaginatedResponse sends a paginated JSON response
func SendPaginatedResponse(w http.ResponseWriter, statusCode int, data interface{}, page, limit int, total int64) {
	SendJSON(w, statusCode, PaginatedResponse{
		Success: true,
		Data:    data,
		Page:    page,
		Limit:   limit,
		Total:   total,
	})
}

func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
