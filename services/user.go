package services

import (
	"rgb/interfaces"
	"rgb/models"

)

type userService struct {
	userRepo interfaces.UserRepository
}

func NewUsersService(repo interfaces.UserRepository) interfaces.UserService {
	return &userService{
		userRepo: repo,
	}
}

func (p *userService) FindAll() ([]models.User, error) {
	return p.userRepo.FindAll()
}

func (p *userService) FindById(id int) (models.User, error) {
	users, err := p.userRepo.FindById(id)
	return users, err
}

func (p *userService) Save(user models.User) (models.User, error) {
	user, err := p.userRepo.Save(user)

	return user, err
}

func (p *userService) Delete(user models.User) error {
	error := p.userRepo.Delete(user)
	return error
}
func (p *userService) AddUser(user models.User) (models.User, error) {
	var err error
	users, err := p.userRepo.Save(user)
	return users, err
}

func (p *userService) Authenticate(email, password string) (models.User, error) {
	users, err := p.userRepo.Authenticate(email, password)
	return users, err
}


