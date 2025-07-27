package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProperty(t *testing.T) {
	router := TestSetup(t)
	defer CleanupTest(t)

	// Get auth token
	token := GetAuthToken(t, router, "owner@example.com", "Password123!")

	tests := []struct {
		name         string
		body         map[string]interface{}
		expectedCode int
		shouldPass   bool
	}{
		{
			name: "Valid property creation",
			body: map[string]interface{}{
				"title":       "Beautiful Apartment",
				"description": "A spacious 2-bedroom apartment in downtown area with modern amenities and great city views.",
				"type":        "apartment",
				"address":     "123 Main St, City, State 12345",
				"price":       1500.0,
				"bedrooms":    2,
				"bathrooms":   1,
				"square_feet": 900,
				"amenities":   []string{"WiFi", "空调", "洗衣机"},
				"available":   true,
			},
			expectedCode: http.StatusCreated,
			shouldPass:   true,
		},
		{
			name: "Missing required fields",
			body: map[string]interface{}{
				"title": "Incomplete Property",
			},
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
		{
			name: "Invalid property type",
			body: map[string]interface{}{
				"title":       "Test Property",
				"description": "A test property with invalid type and sufficient description length to pass validation requirements.",
				"type":        "invalid-type",
				"address":     "123 Test St, Test City, Test State 12345",
				"price":       1000.0,
			},
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
		{
			name: "Negative price",
			body: map[string]interface{}{
				"title":       "Negative Price Property",
				"description": "A test property with negative price and sufficient description length to pass validation requirements.",
				"type":        "apartment",
				"address":     "123 Test St, Test City, Test State 12345",
				"price":       -100.0,
			},
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := map[string]string{
				"Authorization": "Bearer " + token,
			}

			rr := MakeRequest(t, router, "POST", "/api/v1/properties", tt.body, headers)
			assert.Equal(t, tt.expectedCode, rr.Code)

			var response map[string]interface{}
			ParseJSONResponse(t, rr, &response)

			if tt.shouldPass {
				assert.True(t, response["success"].(bool))
				data := response["data"].(map[string]interface{})
				assert.Equal(t, tt.body["title"], data["title"])
				assert.NotNil(t, data["id"])
			} else {
				assert.NotNil(t, response["error"])
			}
		})
	}
}

func TestGetProperties(t *testing.T) {
	router := TestSetup(t)
	defer CleanupTest(t)

	// Get auth token
	token := GetAuthToken(t, router, "user@example.com", "Password123!")

	// Create a test property first
	propertyBody := map[string]interface{}{
		"title":       "Test Property",
		"description": "A test property for listing purposes with sufficient description length for validation.",
		"type":        "apartment",
		"address":     "123 Test St, Test City, Test State 12345",
		"price":       1200.0,
		"bedrooms":    2,
		"bathrooms":   1,
		"available":   true,
	}

	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	rr := MakeRequest(t, router, "POST", "/api/v1/properties", propertyBody, headers)
	assert.Equal(t, http.StatusCreated, rr.Code)

	tests := []struct {
		name         string
		queryParams  string
		expectedCode int
		shouldPass   bool
	}{
		{
			name:         "Get all properties",
			queryParams:  "",
			expectedCode: http.StatusOK,
			shouldPass:   true,
		},
		{
			name:         "Filter by type",
			queryParams:  "?type=apartment",
			expectedCode: http.StatusOK,
			shouldPass:   true,
		},
		{
			name:         "Filter by price range",
			queryParams:  "?min_price=1000&max_price=1500",
			expectedCode: http.StatusOK,
			shouldPass:   true,
		},
		{
			name:         "Pagination",
			queryParams:  "?page=1&limit=10",
			expectedCode: http.StatusOK,
			shouldPass:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/v1/properties" + tt.queryParams
			rr := MakeRequest(t, router, "GET", url, nil, headers)
			assert.Equal(t, tt.expectedCode, rr.Code)

			var response map[string]interface{}
			ParseJSONResponse(t, rr, &response)

			if tt.shouldPass {
				assert.True(t, response["success"].(bool))
				data := response["data"].(map[string]interface{})
				assert.NotNil(t, data["properties"])
				assert.NotNil(t, data["total"])
			} else {
				assert.NotNil(t, response["error"])
			}
		})
	}
}

func TestGetProperty(t *testing.T) {
	router := TestSetup(t)
	defer CleanupTest(t)

	// Get auth token
	token := GetAuthToken(t, router, "user@example.com", "Password123!")

	// Create a test property first
	propertyBody := map[string]interface{}{
		"title":       "Test Property Detail",
		"description": "A detailed test property for getting specific property information with sufficient description length.",
		"type":        "house",
		"address":     "456 Detail St, Detail City, Detail State 67890",
		"price":       2000.0,
		"bedrooms":    3,
		"bathrooms":   2,
		"available":   true,
	}

	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	rr := MakeRequest(t, router, "POST", "/api/v1/properties", propertyBody, headers)
	assert.Equal(t, http.StatusCreated, rr.Code)

	var createResponse map[string]interface{}
	ParseJSONResponse(t, rr, &createResponse)

	data := createResponse["data"].(map[string]interface{})
	propertyID := data["id"].(string)

	tests := []struct {
		name         string
		propertyID   string
		expectedCode int
		shouldPass   bool
	}{
		{
			name:         "Valid property ID",
			propertyID:   propertyID,
			expectedCode: http.StatusOK,
			shouldPass:   true,
		},
		{
			name:         "Invalid property ID format",
			propertyID:   "invalid-id",
			expectedCode: http.StatusBadRequest,
			shouldPass:   false,
		},
		{
			name:         "Non-existent property ID",
			propertyID:   "507f1f77bcf86cd799439011", // Valid ObjectID format but doesn't exist
			expectedCode: http.StatusNotFound,
			shouldPass:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/v1/properties/" + tt.propertyID
			rr := MakeRequest(t, router, "GET", url, nil, headers)
			assert.Equal(t, tt.expectedCode, rr.Code)

			var response map[string]interface{}
			ParseJSONResponse(t, rr, &response)

			if tt.shouldPass {
				assert.True(t, response["success"].(bool))
				data := response["data"].(map[string]interface{})
				assert.Equal(t, propertyBody["title"], data["title"])
				assert.Equal(t, propertyID, data["id"])
			} else {
				assert.NotNil(t, response["error"])
			}
		})
	}
}
