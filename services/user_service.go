package services

import (
	"github.com/uwajacques/go-rest-api/models"
	"github.com/uwajacques/go-rest-api/repositories"
)

type UserService interface {
	GetUser(id string) (models.User, error)
}

type userService struct {
	UserRepository repositories.UserRepository
}

func (us *userService) GetUser(id string) (user models.User, err error) {
	return us.UserRepository.GetUser(id)
}
