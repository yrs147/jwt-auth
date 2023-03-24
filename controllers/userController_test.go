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
