package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/uwajacques/go-rest-api/services"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func (uc *UserController) GetUser(c *gin.Context) {
	userId := ""
	user, err := uc.UserService.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
