package middleware

import (
	"errors"
	"net/http"
	"strings"
)

var (
	// note: supplied via db or env variable for a real world application
	validToken = "adapptor-secure-token"

	ErrNoAuthHeader    = errors.New("no authorization header provided")
	ErrInvalidAuthType = errors.New("invalid authorization type")
	ErrInvalidToken    = errors.New("invalid token")
)

// Auth is a middleware that validates the bearer token from the Authorization header
func Auth(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := validateToken(r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next(w, r)
	})
}

// validates the bearer token from the Authorization header
func validateToken(r *http.Request) error {
	authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
	if authHeader == "" {
		return ErrNoAuthHeader
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ErrInvalidAuthType
	}

	token := parts[1]
	if token != validToken {
		return ErrInvalidToken
	}

	return nil
}

// GetValidToken returns the valid token (for testing purposes)
func GetValidToken() string {
	return validToken
}
