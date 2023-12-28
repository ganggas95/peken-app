package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	_ "gorm.io/driver/postgres"
)

func TestAuthMiddlewareSuccess(t *testing.T) {
	// Initialize
	router, token := InitializeTestApp()

	// Request
	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Read response
	response := recorder.Result()
	// Read response body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	// Status code is 200
	// Response body contains data and value should not to be null
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, responseBody["status"], float64(200))
	assert.NotEqual(t, responseBody["data"], nil)
}

func TestAuthMiddlewareFailed(t *testing.T) {
	// Initialize
	router, _ := InitializeTestApp()

	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Read response
	response := recorder.Result()
	// Read response body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Assertion
	// Status code is 401
	assert.Equal(t, response.StatusCode, http.StatusUnauthorized)
	assert.Equal(t, responseBody["status"], float64(401))
	// Response body contains data and value should be null
	assert.Equal(t, responseBody["data"], nil)
}
