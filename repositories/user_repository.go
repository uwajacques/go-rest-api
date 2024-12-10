package repositories

import "github.com/uwajacques/go-rest-api/models"

type UserRepository interface {
	GetUser(id string) (models.User, error) // Example method to retrieve all users
}

type userRepository struct {
}

func (ur *userRepository) GetUser(id string) (models.User, error) {
	return models.User{
		Id:    "id",
		Email: "jacques@test.dev",
	}, nil
}
