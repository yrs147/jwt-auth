package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yrs147/jwt-auth/controllers"
)

func TestGetUsers(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up the endpoint to test
	router.GET("/users", controllers.GetUsers())

	// Set up a mock context with request parameters
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users?page=1&recordPerPage=5&startIndex=0", nil)
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response body into a map
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	// Check the response contains the expected keys
	assert.Contains(t, response, "total_count")
	assert.Contains(t, response, "user_items")

	// Check the user_items array has the expected length
	userItems := response["user_items"].([]interface{})
	assert.Equal(t, 5, len(userItems))
}

func TestGetUser(t *testing.T) {
	// Setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the user ID param
	c.Params = append(c.Params, gin.Param{Key: "user_id", Value: "user123"})

	// Call the function
	GetUser()(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var user models.User
	if err := json.Unmarshal(w.Body.Bytes(), &user); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}
	if user.UserID != "user123" {
		t.Errorf("Expected user ID %q but got %q", "user123", user.UserID)
	}
}

