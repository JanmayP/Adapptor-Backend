package tests

import (
	"adapptor-backend/pkg/middleware"
	"adapptor-backend/pkg/server"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	t.Run("TestAuthToken", func(t *testing.T) {
		// start server
		server := server.New()

		go func() {
			err := server.Start(8085)
			assert.NoError(t, err)
		}()

		// wait for server to start
		time.Sleep(1 * time.Second)

		t.Run("Empty token", func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://localhost:8085/welcome", nil)
			assert.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		})

		t.Run("Invalid token", func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://localhost:8085/welcome", nil)
			assert.NoError(t, err)

			req.Header.Set("Authorization", "Bearer invalid-token")

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		})

		t.Run("Valid token", func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://localhost:8085/welcome", nil)
			assert.NoError(t, err)

			validToken := middleware.GetValidToken()
			req.Header.Set("Authorization", "Bearer "+validToken)

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})
	})
}
