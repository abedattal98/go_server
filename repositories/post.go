package repositories

import (
	"errors"
	"math/rand"
	"rgb/interfaces"
	"rgb/models"
)

type PostRepository struct {
	db *MemoryStorage
}

func NewPostsRepo(db MemoryStorage) interfaces.IPostRepository {
	return &PostRepository{db: &db}
}
func (p *PostRepository) FindAll() ([]models.Post, error) {
	return p.db.posts, nil
}
func (p *PostRepository) AddPost(userId int, post models.Post) (models.Post, error) {
	post.UserID = userId
	post.ID = rand.Intn(100)

	p.db.posts = append(p.db.posts, post)
	return post, nil
}

func (p *PostRepository) UpdatePost(postID int, post models.Post) (models.Post, error) {
	var err error
	for i, post := range p.db.posts {
		if post.ID == postID {
			p.db.posts[i] = post
			return post, nil
		}
	}
	err = errors.New("Post don't exists")
	return models.Post{}, err
}

func (p *PostRepository) GetPostByID(id int) (models.Post, error) {
	var err error
	for _, p := range p.db.posts {
		if p.ID == id {
			return p, nil
		}
	}
	err = errors.New("Post don't exists")
	return models.Post{}, err
}

func (p *PostRepository) GetPostsByUserID(userId int) []models.Post {
	var posts []models.Post
	for _, p := range p.db.posts {
		if p.UserID == userId {
			posts = append(posts, p)
		}
	}
	return posts
}

func (p *PostRepository) DeletePost(id int) error {
	var err error
	for i, post := range p.db.posts {
		if post.ID == id {
			p.db.posts = append(p.db.posts[:i], p.db.posts[i+1:]...)
			return nil
		}
	}
	err = errors.New("Post don't exists")
	return err
}
