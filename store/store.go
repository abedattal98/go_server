package store

import (
	"errors"
	"rgb/models"
)

var Users []*models.UserEntity
var Posts []*models.PostEntity

func FetchUser(id int) (*models.UserEntity, error) {
	var err error
	user := new(models.UserEntity)
	user.ID = id
	for _, u := range Users {
		if u.ID == user.ID {
			return u, nil
		}
	}
	err = errors.New("User don't exists")
	return nil, err
}
