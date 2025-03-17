package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleWelcome(t *testing.T) {
	req, err := http.NewRequest("GET", "/welcome", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	HandleWelcome(rr, req)

	// assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// unmarshal the response body
	var actual WelcomeResponse
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &actual))

	// assert the response body
	assert.Equal(t, welcomeMessage, actual.Message)
}
