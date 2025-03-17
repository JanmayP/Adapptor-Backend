package api

import (
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendError(w http.ResponseWriter, err error, status int) {
	// return a json error response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}
