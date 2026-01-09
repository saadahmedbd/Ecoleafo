package util

import (
	"encoding/json"
	"net/http"
)

// StandardResponse represents the standard API response structure
type StandardResponse struct {
	Status  string      `json:"status"`  // "success" or "error"
	Message string      `json:"message"` // Human-readable message
	Data    interface{} `json:"data"`    // Response payload (can be nil)
}

// RespondJSON sends a JSON response with standard structure
// Parameters:
//   - w: http.ResponseWriter
//   - statusCode: HTTP status code (200, 201, 400, 404, 500, etc.)
//   - data: Response payload (nil for no data)
//   - message: Human-readable message
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Determine status based on HTTP status code
	status := "success"
	if statusCode >= 400 {
		status = "error"
	}

	// Build response
	response := StandardResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	// Encode and send response
	json.NewEncoder(w).Encode(response)
}

// RespondError sends an error response (shorthand for common errors)
func RespondError(w http.ResponseWriter, statusCode int, message string) {
	RespondJSON(w, statusCode, nil, message)
}

// RespondSuccess sends a success response (shorthand for common success)
func RespondSuccess(w http.ResponseWriter, data interface{}, message string) {
	RespondJSON(w, http.StatusOK, data, message)
}

// RespondCreated sends a 201 Created response
func RespondCreated(w http.ResponseWriter, data interface{}, message string) {
	RespondJSON(w, http.StatusCreated, data, message)
}

// RespondBadRequest sends a 400 Bad Request response
func RespondBadRequest(w http.ResponseWriter, message string) {
	RespondJSON(w, http.StatusBadRequest, nil, message)
}

// RespondUnauthorized sends a 401 Unauthorized response
func RespondUnauthorized(w http.ResponseWriter, message string) {
	RespondJSON(w, http.StatusUnauthorized, nil, message)
}

// RespondForbidden sends a 403 Forbidden response
func RespondForbidden(w http.ResponseWriter, message string) {
	RespondJSON(w, http.StatusForbidden, nil, message)
}

// RespondNotFound sends a 404 Not Found response
func RespondNotFound(w http.ResponseWriter, message string) {
	RespondJSON(w, http.StatusNotFound, nil, message)
}

// RespondInternalError sends a 500 Internal Server Error response
func RespondInternalError(w http.ResponseWriter, message string) {
	RespondJSON(w, http.StatusInternalServerError, nil, message)
}
