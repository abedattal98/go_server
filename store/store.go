package store

import (
	"errors"
	"rgb/models"
)

var Users []*models.User
var Posts []*models.Post

func FetchUser(id int) (*models.User, error) {
	var err error
	user := new(models.User)
	user.ID = id
	for _, u := range Users {
		if u.ID == user.ID {
			return u, nil
		}
	}
	err = errors.New("Current context user not set")
	return nil, err
}
