package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleAction(t *testing.T) {
	t.Run("empty request body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/action", strings.NewReader(""))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		HandleAction(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("invalid request body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/action", strings.NewReader(`{"action": "invalid"}`))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		HandleAction(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("primary action", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/action", strings.NewReader(`{"action": "primary"}`))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		HandleAction(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response ActionResponse
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "Action processed successfully", response.Message)

	})

	t.Run("secondary action", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/action", strings.NewReader(`{"action": "secondary"}`))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		HandleAction(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response ActionResponse
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "Secondary action processed successfully", response.Message)
	})

	t.Run("invalid action", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/action", strings.NewReader(`{"action": "invalid"}`))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		HandleAction(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response ErrorResponse
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "invalid action", response.Error)
	})
}
