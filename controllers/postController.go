package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhil/blogging-platform-api/config"
	"github.com/nikhil/blogging-platform-api/models"
)

// Create a new post
func CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Create(&post)
    c.JSON(http.StatusCreated, post)
}

// Get all posts
// GetAllPosts (updated for search functionality)
func GetAllPosts(c *gin.Context) {
    var posts []models.Post
    term := c.DefaultQuery("term", "") // Get the 'term' query parameter

    if term != "" {
        // Search for the term in title, content, and category
        config.DB.Where("title ILIKE ? OR content ILIKE ? OR category ILIKE ?", "%"+term+"%", "%"+term+"%", "%"+term+"%").Find(&posts)
    } else {
        // If no term is provided, return all posts
        config.DB.Find(&posts)
    }

    c.JSON(http.StatusOK, posts)
}


// Get a post by ID
func GetPostByID(c *gin.Context) {
    id := c.Param("id")
    var post models.Post

    if err := config.DB.First(&post, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    c.JSON(http.StatusOK, post)
}

// Update a post
func UpdatePost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post

    if err := config.DB.First(&post, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    var input models.Post
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    post.Title = input.Title
    post.Content = input.Content
    post.Category = input.Category
    post.Tags = input.Tags

    config.DB.Save(&post)
    c.JSON(http.StatusOK, post)
}

// Delete a post
func DeletePost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post

    if err := config.DB.First(&post, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    config.DB.Delete(&post)
    c.Status(http.StatusNoContent)
}
