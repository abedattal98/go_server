package services

import (
	"errors"
	"rgb/models"
	"rgb/store"
	"math/rand"
)

func AddPost(userId int, post *models.Post) error {
	post.UserID = userId
	post.ID = rand.Intn(100);

	store.Posts = append(store.Posts, post)
		return nil
}

func UpdatePost(postID int, post *models.Post) (error) {
	var err error
	for i, p := range store.Posts {
		if p.ID == postID  {
			store.Posts[i] = post
			return nil
		}
	}
	err = errors.New("Post don't exists")
	return err
}

func GetPostByID(id int) (*models.Post, error) {
	var err error
	for _, p := range store.Posts {
		if p.ID == id  {
			return p, nil
		}
	}
	err = errors.New("Post don't exists")
	return nil, err
}

func GetPostsByUserID(userId int) []*models.Post {
	var posts []*models.Post
	for _, p := range store.Posts {
		if p.UserID == userId {
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
