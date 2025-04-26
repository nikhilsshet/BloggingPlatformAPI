package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhil/blogging-platform-api/config"
	"github.com/nikhil/blogging-platform-api/routes"
)

func main() {
    // Initialize database
    config.ConnectDatabase()

    // Set up router
    router := gin.Default()

    // Register routes
    routes.RegisterPostRoutes(router)

    // Run the server
    router.Run(":8080")
}
