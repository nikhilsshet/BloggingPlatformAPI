package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhil/blogging-platform-api/controllers"
)

func RegisterPostRoutes(router *gin.Engine) {
    posts := router.Group("/posts")
    {
        posts.POST("/", controllers.CreatePost)      // Create a post
        posts.GET("/", controllers.GetAllPosts)      // Get all posts
        posts.GET("/:id", controllers.GetPostByID)   // Get a post by ID
        posts.PUT("/:id", controllers.UpdatePost)    // Update a post by ID
        posts.DELETE("/:id", controllers.DeletePost) // Delete a post by ID
    }
}
