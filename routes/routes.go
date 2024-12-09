package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/uwajacques/go-rest-api/middlewares"
    "net/http"
)

func SetupRoutes(router *gin.Engine) {

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
            userGroup.GET("/", func(c *gin.Context) {
                c.JSON(http.StatusOK, "Get user")
            })

            userGroup.POST("/", func(c *gin.Context) {
                c.JSON(http.StatusOK, "Create user")
            })
        }
    }

}
