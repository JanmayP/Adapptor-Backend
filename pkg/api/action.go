package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	PrimaryAction   string = "primary"
	SecondaryAction string = "secondary"
)

// ActionRequest represents the request body for the action endpoint
type ActionRequest struct {
	Action string `json:"action" example:"primary" enums:"primary,secondary"` // The action to perform
}

// ActionResponse represents the response from the action endpoint
type ActionResponse struct {
	Status  string `json:"status" example:"success"`           // The status of the action
	Message string `json:"message" example:"Action processed"` // A descriptive message
}

// @Summary Process an action
// @Description Process one of the available actions (primary or secondary)
// @Tags action
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body ActionRequest true "Action request"
// @Success 200 {object} ActionResponse
// @Failure 400 {object} ErrorResponse "Invalid request body or action"
// @Failure 401 {object} ErrorResponse "Unauthorized - Invalid or missing token"
// @Example {json} Example Request - Primary Action:
//
//	{
//	  "action": "primary"
//	}
//
// @Example {json} Example Request - Secondary Action:
//
//	{
//	  "action": "secondary"
//	}
//
// @Example {json} Success Response:
//
//	{
//	  "status": "success",
//	  "message": "Action processed successfully"
//	}
//
// @Router /action [post]
func HandleAction(w http.ResponseWriter, r *http.Request) {
	var request ActionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError(w, errors.New("invalid request body"), http.StatusBadRequest)
		return
	}

	var response ActionResponse

	switch request.Action {
	case PrimaryAction:
		response = ActionResponse{
			Status:  "success",
			Message: "Action processed successfully",
		}
	case SecondaryAction:
		response = ActionResponse{
			Status:  "success",
			Message: "Secondary action processed successfully",
		}
	default:
		sendError(w, errors.New("invalid action"), http.StatusBadRequest)
		return
	}

	sendResponse(w, response)
}
