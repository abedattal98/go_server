package domain

import (
	"rgb/models"
)


type IUserRepository interface {
	FindAll() ([]models.User, error)
	FindByID( id int) (models.User, error)
	Save( user models.User) (models.User, error)
	Delete(user models.User) error
}