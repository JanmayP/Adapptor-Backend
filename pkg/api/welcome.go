package api

import (
	"net/http"
)

const welcomeMessage = "Welcome to the Adapptor Backend API!"

type WelcomeResponse struct {
	Message string `json:"message" example:"Welcome to the Adapptor Backend API!"` // Welcome message
}

// @Summary Get welcome message
// @Description Returns a welcome message
// @Tags welcome
// @Produce json
// @Security BearerAuth
// @Success 200 {object} WelcomeResponse
// @Failure 401 {object} ErrorResponse "Unauthorized - Invalid or missing token"
// @Example {json} Success Response:
//
//	{
//	  "message": "Welcome to the Adapptor Backend API!"
//	}
//
// @Router /welcome [get]
func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	response := WelcomeResponse{
		Message: welcomeMessage,
	}
	sendResponse(w, response)
}
