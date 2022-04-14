package services

import (
	"errors"
	"rgb/models"
	"rgb/store"
	"time"
)

func AddUser(user *models.User) (*models.UserEntity,error) {
	var err error

	for _, u := range store.Users {
		if (u.Username == user.Username){
			err = errors.New("User already exists")
			return nil,err
		}
	}
	userEntity := &models.UserEntity{
		Entity: models.Entity{
			ID:        int(time.Now().Unix()),
			CreatedAt: time.Date(2020, time.July, 30, 0, 0, 0, 0, time.UTC),
		},
		User: *user,
	}
	store.Users = append(store.Users, userEntity)
		return userEntity,nil
}

func Authenticate(username, password string) (*models.UserEntity, error) {
	var err error;
	for _, u := range store.Users {
		if (u.Username == username && u.Password == password ){
			return u, nil
		}
	}
	err = errors.New("User don't exists")
	return nil,err
}
