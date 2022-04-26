package interfaces

import "rgb/models"

type IPostService interface {
	UpdatePost(postID int, post models.Post) (models.Post, error)
	GetPostByID(id int) (models.Post, error)
	GetPostsByUserID(userId int) []models.Post
	DeletePost(id int) error
	AddPost(userId int, post models.Post) (models.Post, error)
}
type IPostRepository interface {
	UpdatePost(postID int, post models.Post) (models.Post, error)
	GetPostByID(id int) (models.Post, error)
	GetPostsByUserID(userId int) []models.Post
	DeletePost(id int) error
	AddPost(userId int, post models.Post) (models.Post, error)
}
