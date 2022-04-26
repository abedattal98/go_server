package repositories

import (
	"errors"
	"rgb/interfaces"
	"rgb/models"
)

type UserRepository struct {
	db *MemoryStorage
}

func NewUsersRepo(db MemoryStorage) interfaces.IUserRepository {
	return &UserRepository{db: &db}
}
func (p *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	users = p.db.users
	return users, error(nil)
}

func (p *UserRepository) FindByID(id int) (models.User, error) {
	var user models.User
	for _, u := range p.db.users {
		if u.ID == id {
			return u, error(nil)
		}
	}
	return user, error(nil)
}

func (p *UserRepository) Save(user models.User) (models.User, error) {
	for _, u := range p.db.users {
		if u.Email == user.Email {
			err := errors.New("User already exists")
			return models.User{}, err
		}
	}
	p.db.users = append(p.db.users, user)
	return user, nil
}

func (p *UserRepository) Delete(user models.User) error {
	for i, u := range p.db.users {
		if u.ID == user.ID {
			p.db.users = append(p.db.users[:i], p.db.users[i+1:]...)
		}
	}
	return nil
}

func (p *UserRepository) Authenticate(email, password string) (models.User, error) {

	for _, u := range p.db.users {
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
	for _, u := range p.db.users {
		if u.ID == user.ID {
			return u, nil
		}
	}
	err = errors.New("User don't exists")
	return *user, err
}
