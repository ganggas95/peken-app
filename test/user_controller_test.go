package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreateUserAPISuccess(t *testing.T) {
	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(
		`{
			"email": "user2@test.com",
			"name": "Test",
			"username": "test2",
			"password": "password",
			"roles": [{"id": 1 }]
		}`)
	request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
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
	// Status code is 201
	assert.Equal(t, response.StatusCode, http.StatusCreated)
	assert.Equal(t, responseBody["status"], float64(201))

	assert.Equal(t, responseBody["data"].(map[string]interface{})["email"], "user2@test.com")
	assert.NotEqual(t, responseBody["data"].(map[string]interface{})["id"], nil)
}

func TestCreateUserAPIFailed(t *testing.T) {
	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(
		`{
			"email": "",
			"name": "",
			"username": "",
			"password": ""
		}`)
	request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
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
	// Status code is 400
	assert.Equal(t, response.StatusCode, http.StatusBadRequest)
	assert.Equal(t, responseBody["status"], float64(400))

	assert.Equal(t, responseBody["data"], nil)
}

func TestCreateUserAPIConflict(t *testing.T) {
	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(
		`{
			"email": "test@test.com",
			"name": "test",
			"username": "testuser",
			"password": "password",
			"roles": [{"id": 1}]
		}`)
	request := httptest.NewRequest(http.MethodPost, "/api/users", requestBody)
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
	// Status code is 409
	assert.Equal(t, response.StatusCode, http.StatusConflict)
	assert.Equal(t, responseBody["status"], float64(409))

	assert.Equal(t, responseBody["data"], nil)
}

func TestGetDetailUserAPISuccess(t *testing.T) {
	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "/api/users/1", requestBody)
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
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, responseBody["status"], float64(200))

	assert.NotEqual(t, responseBody["data"], nil)
}
func TestGetDetailUserAPINotFound(t *testing.T) {
	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "/api/users/50", requestBody)
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
	// Status code is 404
	assert.Equal(t, response.StatusCode, http.StatusNotFound)
	assert.Equal(t, responseBody["status"], float64(404))

	assert.Equal(t, responseBody["data"], nil)
}

func TestUpdateDetailUserAPISuccess(t *testing.T) {

	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(
		`{
			"email": "user2@test.com",
			"name": "Test",
			"username": "test2",
			"roles": [{"id": 1 }]
		}`)
	request := httptest.NewRequest(http.MethodPut, "/api/users/1", requestBody)
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
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, responseBody["status"], float64(200))

	assert.Equal(t, responseBody["data"].(map[string]interface{})["email"], "user2@test.com")
	assert.NotEqual(t, responseBody["data"].(map[string]interface{})["id"], nil)
}

func TestUpdateDetailUserAPIFailed(t *testing.T) {

	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(
		`{
			"email": "user2@test.com",
			"name": "Test",
			"username": "test2",
			"roles": [{"id": 1 }]
		}`)
	request := httptest.NewRequest(http.MethodPut, "/api/users/30", requestBody)
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
	// Status code is 404
	assert.Equal(t, response.StatusCode, http.StatusNotFound)
	assert.Equal(t, responseBody["status"], float64(404))

	assert.Equal(t, responseBody["data"], nil)
}

func TestUpdateDetailUserAPIConflict(t *testing.T) {

	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(
		`{
			"email": "test@test.com",
			"name": "Test",
			"username": "testuser",
			"roles": [{"id": 1 }]
		}`)
	request := httptest.NewRequest(http.MethodPut, "/api/users/2", requestBody)
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
	// Status code is 409
	assert.Equal(t, response.StatusCode, http.StatusConflict)
	assert.Equal(t, responseBody["status"], float64(409))

	assert.Equal(t, responseBody["data"], nil)
}

func TestDeleteDetailUserAPISuccess(t *testing.T) {

	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodDelete, "/api/users/1", requestBody)
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
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, responseBody["status"], float64(200))

	assert.Equal(t, responseBody["data"], nil)

}

func TestDeleteDetailUserAPIFailed(t *testing.T) {

	router, token := InitializeTestApp()

	// Request body
	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodDelete, "/api/users/30", requestBody)
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
	// Status code is 404
	assert.Equal(t, response.StatusCode, http.StatusNotFound)
	assert.Equal(t, responseBody["status"], float64(404))

	assert.Equal(t, responseBody["data"], nil)

}

func TestGetListUsersAPISuccess(t *testing.T) {

	router, token := InitializeTestApp()

	// Request body
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
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, responseBody["status"], float64(200))

	assert.NotEqual(t, responseBody["data"], nil)
}
