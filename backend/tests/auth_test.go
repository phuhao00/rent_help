package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegistration(t *testing.T) {
	router := TestSetup(t)
	defer CleanupTest(t)

	tests := []struct {
		name         string
		body         map[string]interface{}
		expectedCode int
		shouldPass   bool
	}{
		{
			name: "Valid registration",
			body: map[string]interface{}{
				"email":     "test@example.com",
				"password":  "Password123!",
				"full_name": "Test User",
				"phone":     "+1234567890",
			},
			expectedCode: http.StatusCreated,
			shouldPass:   true,
		},
		{
			name: "Invalid email",
			body: map[string]interface{}{
				"email":     "invalid-email",
				"password":  "Password123!",
				"full_name": "Test User",
			},
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
		{
			name: "Weak password",
			body: map[string]interface{}{
				"email":     "test2@example.com",
				"password":  "weak",
				"full_name": "Test User",
			},
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
		{
			name: "Missing required fields",
			body: map[string]interface{}{
				"email": "test3@example.com",
			},
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
		{
			name: "Duplicate email",
			body: map[string]interface{}{
				"email":     "test@example.com", // Same as first test
				"password":  "Password123!",
				"full_name": "Another User",
			},
			expectedCode: http.StatusConflict,
			shouldPass:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := MakeRequest(t, router, "POST", "/api/v1/auth/register", tt.body, nil)
			assert.Equal(t, tt.expectedCode, rr.Code)

			var response map[string]interface{}
			ParseJSONResponse(t, rr, &response)

			if tt.shouldPass {
				assert.True(t, response["success"].(bool))
				assert.NotNil(t, response["data"])
			} else {
				assert.NotNil(t, response["error"])
			}
		})
	}
}

func TestUserLogin(t *testing.T) {
	router := TestSetup(t)
	defer CleanupTest(t)

	// First, register a user
	registerBody := map[string]interface{}{
		"email":     "test@example.com",
		"password":  "Password123!",
		"full_name": "Test User",
	}
	rr := MakeRequest(t, router, "POST", "/api/v1/auth/register", registerBody, nil)
	assert.Equal(t, http.StatusCreated, rr.Code)

	tests := []struct {
		name         string
		body         map[string]interface{}
		expectedCode int
		shouldPass   bool
	}{
		{
			name: "Valid login",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"password": "Password123!",
			},
			expectedCode: http.StatusOK,
			shouldPass:   true,
		},
		{
			name: "Invalid password",
			body: map[string]interface{}{
				"email":    "test@example.com",
				"password": "wrongpassword",
			},
			expectedCode: http.StatusUnauthorized,
			shouldPass:   false,
		},
		{
			name: "Non-existent user",
			body: map[string]interface{}{
				"email":    "nonexistent@example.com",
				"password": "Password123!",
			},
			expectedCode: http.StatusUnauthorized,
			shouldPass:   false,
		},
		{
			name: "Missing fields",
			body: map[string]interface{}{
				"email": "test@example.com",
			},
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := MakeRequest(t, router, "POST", "/api/v1/auth/login", tt.body, nil)
			assert.Equal(t, tt.expectedCode, rr.Code)

			var response map[string]interface{}
			ParseJSONResponse(t, rr, &response)

			if tt.shouldPass {
				assert.True(t, response["success"].(bool))
				data := response["data"].(map[string]interface{})
				assert.NotNil(t, data["access_token"])
				assert.NotNil(t, data["user"])
			} else {
				assert.NotNil(t, response["error"])
			}
		})
	}
}

func TestGetProfile(t *testing.T) {
	router := TestSetup(t)
	defer CleanupTest(t)

	// Get auth token
	token := GetAuthToken(t, router, "test@example.com", "Password123!")

	tests := []struct {
		name         string
		token        string
		expectedCode int
		shouldPass   bool
	}{
		{
			name:         "Valid request with token",
			token:        token,
			expectedCode: http.StatusOK,
			shouldPass:   true,
		},
		{
			name:         "Request without token",
			token:        "",
			expectedCode: http.StatusUnauthorized,
			shouldPass:   false,
		},
		{
			name:         "Request with invalid token",
			token:        "invalid-token",
			expectedCode: http.StatusUnauthorized,
			shouldPass:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := make(map[string]string)
			if tt.token != "" {
				headers["Authorization"] = "Bearer " + tt.token
			}

			rr := MakeRequest(t, router, "GET", "/api/v1/users/profile", nil, headers)
			assert.Equal(t, tt.expectedCode, rr.Code)

			var response map[string]interface{}
			ParseJSONResponse(t, rr, &response)

			if tt.shouldPass {
				assert.True(t, response["success"].(bool))
				data := response["data"].(map[string]interface{})
				assert.Equal(t, "test@example.com", data["email"])
			} else {
				assert.NotNil(t, response["error"])
			}
		})
	}
}
