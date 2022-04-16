package repositories

import (
	"rgb/models"
	services "rgb/repositories/interface"
)

var UserStore []models.User

type UserRepository struct {
}

func ProvideUserRepository() services.IUserService {
	return &UserRepository{}
}

func (p *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	users = UserStore
	return users, error(nil)
}

func (p *UserRepository) FindByID(id int) (models.User, error) {
	var user models.User
	for _, u := range UserStore {
		if u.ID == id {
			return u, error(nil)
		}
	}
	return user, error(nil)
}

func (p *UserRepository) Save(user models.User) (models.User, error) {
	UserStore = append(UserStore, user)
	return user, nil
}

func (p *UserRepository) Delete(user models.User) error {
	for i, u := range UserStore {
		if u.ID == user.ID {
			UserStore = append(UserStore[:i], UserStore[i+1:]...)
		}
	}
	return nil
}
