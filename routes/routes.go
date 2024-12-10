package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uwajacques/go-rest-api/controllers"
	"github.com/uwajacques/go-rest-api/middlewares"
	"github.com/uwajacques/go-rest-api/repositories"
	"github.com/uwajacques/go-rest-api/services"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	// db := setupDatabase()
	userRepo := repositories.UserRepository()
	userService := services.UserService(userRepo)
	userController := &controllers.UserController{UserService: userService}

	api := router.Group("/api")
	{
		// Auth routes
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", func(c *gin.Context) {
				c.JSON(http.StatusOK, "Logged in")
			})

			authGroup.POST("/logout", func(c *gin.Context) {
				c.JSON(http.StatusOK, "Logged out")
			})

			authGroup.POST("/reset-password", func(c *gin.Context) {
				c.JSON(http.StatusOK, "Password Reset")
			})
		}

		// User routes
		userGroup := api.Group("/user")
		userGroup.Use(middlewares.AuthMiddleware())
		{
			userGroup.GET("/", userController.GetUser)

			userGroup.POST("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, "Create user")
			})
		}
	}

}
