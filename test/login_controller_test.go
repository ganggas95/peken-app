package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"peken-be/app"
	"peken-be/helper"
	"peken-be/models/domain"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
)

func Init() {
	// app.InitLog()
	godotenv.Load("../.env")
	os.Setenv("ENV", "test")
}

func InitializeTestApp() *gin.Engine {
	Init()
	// Initialize db and router
	db := app.ConnectToDb()
	router := InitRouter(db)

	// Clean up database and make migration
	db.Migrator().DropTable(&domain.User{})
	db.AutoMigrate(&domain.User{})
	// Initializer User data
	passwordUtils := helper.NewPasswordUtils()
	password, _ := passwordUtils.HashPassword("password")
	db.Create(&domain.User{
		Username: "testuser",
		Password: password,
	})
	return router
}

func TestLoginSuccess(t *testing.T) {
	// Initialize
	router := InitializeTestApp()

	// Request
	requestBody := strings.NewReader(`{"username": "testuser", "password": "password"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/login", requestBody)
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
	// Status code is 200
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, responseBody["status"], float64(200))
	// Response body contains access_token and value is not null
	assert.NotEqual(t, responseBody["data"].(map[string]interface{})["access_token"], nil)
}

func TestLoginFailed(t *testing.T) {
	// Initialize
	router := InitializeTestApp()

	requestBody := strings.NewReader(`{"username": "testwrongpass", "password": "password"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/login", requestBody)
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
	// Response body contains access_token and value is not null
	assert.Equal(t, responseBody["data"], nil)
}
