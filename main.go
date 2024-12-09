package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uwajacques/go-rest-api/routes"
)

func main() {
	// Load configuration
	//	config.LoadEnv()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
