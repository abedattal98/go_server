package services

import (
	"errors"
	"rgb/models"
	"rgb/store"
	"time"
)

func AddUser(user *models.User) (*models.User,error) {
	var err error

	for _, u := range store.Users {
		if (u.Username == user.Username){
			err = errors.New("User already exists")
			return nil,err
		}
	}
	user.ID  = int(time.Now().Unix());
	store.Users = append(store.Users, user)
		return user,nil
}

func Authenticate(username, password string) (*models.User, error) {
	var err error;

	user := new(models.User)
	for _, u := range store.Users {
		if (u.Username == user.Username && u.Password == user.Password ){
			return u, nil
		}
	}
	err = errors.New("User don't exists")
	return nil,err
}
