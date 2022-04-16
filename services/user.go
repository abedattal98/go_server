package services

import (
	"errors"
	"rgb/domain"
	"rgb/models"
	services "rgb/repositories/interface"
	"rgb/store"
	"time"
)

type UserService struct {
	userRepo domain.IUserRepository
}

func ProvideUserService(repo domain.IUserRepository) services.IUserService {
	return &UserService{
		userRepo: repo,
	}
}

func (p *UserService) FindAll() ([]models.User, error) {
	return p.userRepo.FindAll()
}

func (p *UserService) FindByID(id int) (models.User, error) {
	users, err := p.userRepo.FindByID(id)
	return users, err
}

func (p *UserService) Save(user models.User) (models.User, error) {
	user, err := p.userRepo.Save(user)

	return user, err
}

func (p *UserService) Delete(user models.User) error {
	error := p.userRepo.Delete(user)
	return error
}

func AddUser(user *models.User) (*models.User, error) {
	var err error

	for _, u := range store.Users {
		if u.Username == user.Username {
			err = errors.New("User already exists")
			return nil, err
		}
	}
	user.ID = int(time.Now().Unix())
	store.Users = append(store.Users, user)
	return user, nil
}

func Authenticate(username, password string) (*models.User, error) {
	var err error

	for _, u := range store.Users {
		if u.Username == username && u.Password == password {
			return u, nil
		}
	}
	err = errors.New("User don't exists")
	return nil, err
}
