package repositories

import (
	"rgb/models"
)

var UserStore []models.User

type UserRepository struct {
}

func ProvideUserRepository() UserRepository {
	return UserRepository{}
}

func (p *UserRepository) FindAll() []models.User {
	var users []models.User
	users = UserStore
	return users
}

func (p *UserRepository) FindById(id int) models.User {
	var user models.User
	for _, u := range UserStore {
		if u.ID == id  {
			return u
		}
	}
	return user
}

func (p *UserRepository) Save(user models.User) models.User {
	UserStore = append(UserStore, user)
	return user
}

func (p *UserRepository) Delete(user models.User) error {
	for i, u := range UserStore {
		if u.ID == user.ID {
			UserStore = append(UserStore[:i], UserStore[i+1:]...)
		}
	}
	return nil
}