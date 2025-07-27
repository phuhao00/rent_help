package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"rent-help-backend/internal/handlers"
	"rent-help-backend/internal/middleware"
	"rent-help-backend/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestSetup sets up the test environment
func TestSetup(t *testing.T) *gin.Engine {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize test database
	if err := database.InitTestDatabase(); err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	// Create router
	router := gin.New()

	// Add middleware
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorHandler())

	// Setup routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/profile", handlers.GetProfile)
				users.PUT("/profile", handlers.UpdateProfile)
			}

			// Property routes
			properties := protected.Group("/properties")
			{
				properties.POST("/", handlers.CreateProperty)
				properties.GET("/", handlers.GetProperties)
				properties.GET("/:id", handlers.GetProperty)
				properties.PUT("/:id", handlers.UpdateProperty)
				properties.DELETE("/:id", handlers.DeleteProperty)
			}

			// Booking routes
			bookings := protected.Group("/bookings")
			{
				bookings.POST("/", handlers.CreateBooking)
				bookings.GET("/", handlers.GetUserBookings)
				bookings.GET("/:id", handlers.GetBooking)
				bookings.PUT("/:id/status", handlers.UpdateBookingStatus)
			}
		}
	}

	return router
}

// MakeRequest makes an HTTP request for testing
func MakeRequest(t *testing.T, router *gin.Engine, method, url string, body interface{}, headers map[string]string) *httptest.ResponseRecorder {
	var reqBody *bytes.Buffer

	if body != nil {
		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)
		reqBody = bytes.NewBuffer(jsonBody)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequest(method, url, reqBody)
	assert.NoError(t, err)

	// Set default headers
	req.Header.Set("Content-Type", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create response recorder
	rr := httptest.NewRecorder()

	// Perform request
	router.ServeHTTP(rr, req)

	return rr
}

// ParseJSONResponse parses JSON response body
func ParseJSONResponse(t *testing.T, rr *httptest.ResponseRecorder, target interface{}) {
	err := json.Unmarshal(rr.Body.Bytes(), target)
	assert.NoError(t, err)
}

// GetAuthToken gets an authentication token for testing
func GetAuthToken(t *testing.T, router *gin.Engine, email, password string) string {
	// First register the user
	registerBody := map[string]interface{}{
		"email":     email,
		"password":  password,
		"full_name": "Test User",
	}

	rr := MakeRequest(t, router, "POST", "/api/v1/auth/register", registerBody, nil)
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Then login to get token
	loginBody := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	rr = MakeRequest(t, router, "POST", "/api/v1/auth/login", loginBody, nil)
	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	ParseJSONResponse(t, rr, &response)

	data, ok := response["data"].(map[string]interface{})
	assert.True(t, ok)

	token, ok := data["access_token"].(string)
	assert.True(t, ok)

	return token
}

// CleanupTest cleans up after tests
func CleanupTest(t *testing.T) {
	if err := database.CleanTestDatabase(); err != nil {
		t.Errorf("Failed to clean test database: %v", err)
	}
}
