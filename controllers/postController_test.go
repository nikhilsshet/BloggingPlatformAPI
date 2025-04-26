package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nikhil/blogging-platform-api/config"
)

func TestCreatePost(t *testing.T) {
    router := gin.Default()
    config.ConnectDatabase()

    // Define the endpoint
    router.POST("/posts", CreatePost)

    // Create a new post request
    newPost := `{
        "title": "Test Post",
        "content": "Content for the test post."
    }`

    req, err := http.NewRequest("POST", "/posts", bytes.NewBufferString(newPost))
    if err != nil {
        t.Fatal(err)
    }

    // Record the response
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Check if the response status code is 201
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status code %d but got %d", http.StatusCreated, w.Code)
    }

    // You can also check the response body or database state here if needed.
}
