package services

import (
	"errors"
	"math/rand"
	"rgb/models"
	"rgb/store"
)

func AddPost(userId int, post *models.PostEntity) error {
	var err error
	post.UserID = userId
	post.ID = rand.Intn(100)

	if post.UserID != 0 {
		store.Posts = append(store.Posts, post)
		return nil
	}

	return err
}

func UpdatePost(postID int, post *models.PostEntity) error {
	var err error
	for i, p := range store.Posts {
		if p.ID == postID {
			store.Posts[i] = post
			return nil
		}
	}
	err = errors.New("Post don't exists")
	return err
}

func GetPostByID(id int) (*models.PostEntity, error) {
	var err error
	for _, p := range store.Posts {
		if p.ID == id {
			return p, nil
		}
	}
	err = errors.New("Post don't exists")
	return nil, err
}

func GetPostsByUserID(user *models.UserEntity) []*models.PostEntity {
	var posts []*models.PostEntity
	for _, p := range store.Posts {
		if p.UserID == user.ID {
			posts = append(posts, p)
		}
	}
	return posts
}

func DeletePost(id int) error {
	var err error
	for i, p := range store.Posts {
		if p.ID == id {
			store.Posts = append(store.Posts[:i], store.Posts[i+1:]...)
			return nil
		}
	}
	err = errors.New("Post don't exists")
	return err
}
