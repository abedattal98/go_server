package services

import (
	"rgb/interfaces"
	"rgb/models"
)

type postService struct {
	postRepo interfaces.PostRepository
}

func NewPostsService(repo interfaces.PostRepository) interfaces.PostService {
	return &postService{
		postRepo: repo,
	}
}

func (p *postService) GetPostsByUserID(userId int) []models.Post {
	return p.postRepo.GetPostsByUserID(userId)
}

func (p *postService) GetPostByID(id int) (models.Post, error) {
	return p.postRepo.GetPostByID(id)
}

func (p *postService) AddPost(userID int, post models.Post) (models.Post, error) {
	return p.postRepo.AddPost(post.UserID, post)
}

func (p *postService) DeletePost(id int) error {
	error := p.postRepo.DeletePost(id)
	return error
}

func (p *postService) UpdatePost(id int, post models.Post) (models.Post, error) {
	return p.postRepo.UpdatePost(id, post)
}
