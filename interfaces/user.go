package interfaces

import (
	"rgb/models"
)

type UserService interface {
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	Save(user models.User) (models.User, error)
	Delete(user models.User) error
	Authenticate(email, password string) (models.User, error)
}
type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	Save(user models.User) (models.User, error)
	Delete(user models.User) error
	Authenticate(username, password string) (models.User, error)
	FetchUser(id int) (models.User, error)
}
