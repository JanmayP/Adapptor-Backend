package api

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request body"` // Error message
}
