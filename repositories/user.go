package repositories

import (
	"errors"
	"rgb/domain"
	"rgb/models"
)

var UserStore = []models.User{
	models.User{
		ID:       1,
		Email:    "admin@admin.com",
		Password: "admin",
		Username: "admin",
	},
}

type UserRepository struct {
}

func ProvideUserRepository() domain.IUserRepository {
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
	for _, u := range UserStore {
		if u.Email == user.Email {
			err := errors.New("User already exists")
			return models.User{}, err
		}
	}
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

func (p *UserRepository) Authenticate(email, password string) (models.User, error) {

	for _, u := range UserStore {
		if u.Email == email && u.Password == password {
			return u, nil
		}
	}
	var user models.User = models.User{}
	err := errors.New("User don't exists")
	return user, err
}

func (p *UserRepository) FetchUser(id int) (models.User, error) {
	var err error
	user := new(models.User)
	user.ID = id
	for _, u := range UserStore {
		if u.ID == user.ID {
			return u, nil
		}
	}
	err = errors.New("User don't exists")
	return *user, err
}
