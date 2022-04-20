package services

import (
	"log"
	"rgb/domain"
	"rgb/models"
	"rgb/repositories"
	"rgb/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addTestPost(repo domain.IUserRepository) (models.Post, error) {
	user, err := repositories.AddTestUser(repo)

	post := models.Post{
		Title:   "Gotham cronicles",
		Content: "Joker is planning big hit tonight.",
		UserID: user.ID,
	}
	log.Println(1111)
	log.Println(user)
	postAdded, err := AddPost(user.ID, post)
	log.Println(postAdded)
	return postAdded, err
}

func TestAddPost(t *testing.T) {
	store.Posts = nil
	userRepository := repositories.ProvideUserRepository()
	post, err := addTestPost(userRepository)
	assert.NoError(t, err)
	assert.Greater(t, post.ID, 0)
}

func TestGetPostsByUserID(t *testing.T) {
	userRepository := repositories.ProvideUserRepository()
	post, err := addTestPost(userRepository)
	assert.NoError(t, err)
	posts := GetPostsByUserID(post.UserID)
	assert.NoError(t, err)
	assert.Equal(t, post, posts[0])
}

func TestGetPostsByUserIDEmpty(t *testing.T) {
	userRepository := repositories.ProvideUserRepository()
	user, err := repositories.AddTestUser(userRepository)
	assert.NoError(t, err)
	posts := GetPostsByUserID(user.ID)
	assert.NoError(t, err)
	assert.Empty(t, posts)
}

func TestFetchPost(t *testing.T) {
	userRepository := repositories.ProvideUserRepository()
	post, err := addTestPost(userRepository)
	assert.NoError(t, err)

	fetchedPost, err := GetPostByID(post.ID)
	assert.NoError(t, err)
	assert.Equal(t, post.ID, fetchedPost.ID)
	assert.Equal(t, post.Title, fetchedPost.Title)
	assert.Equal(t, post.Content, fetchedPost.Content)
	assert.Equal(t, post.UserID, fetchedPost.UserID)
}

func TestFetchNotExistingPost(t *testing.T) {
	fetchedPost, err := GetPostByID(0)
	assert.Error(t, err)
	assert.Empty(t, fetchedPost)
	assert.Equal(t, "Post don't exists", err.Error())
}

func TestUpdatePost(t *testing.T) {
	userRepository := repositories.ProvideUserRepository()
	post, err := addTestPost(userRepository)
	assert.NoError(t, err)

	post.Title = "New title"
	post.Content = "New content"
	post,err = UpdatePost(post.ID, post)
	assert.NoError(t, err)
}

func TestDeletePost(t *testing.T) {
	userRepository := repositories.ProvideUserRepository()
	post, err := addTestPost(userRepository)
	assert.NoError(t, err)
	err = DeletePost(post.ID)
	assert.NoError(t, err)
}
