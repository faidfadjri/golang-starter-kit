package response

import (
	"encoding/json"
	"net/http"
)

// APIResponse adalah template untuk semua response JSON
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewSuccessResponse membuat response sukses
func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse membuat response error
func NewErrorResponse(message string, err error) APIResponse {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	return APIResponse{
		Success: false,
		Message: message,
		Error:   errorMessage,
	}
}

// JSONResponse mengirim response sebagai JSON
func JSONResponse(w http.ResponseWriter, statusCode int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
