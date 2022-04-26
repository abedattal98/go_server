package repositories

// import (
// 	"testing"
// 	"rgb/services"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAddPost(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)

// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)
// 	assert.Greater(t, post.ID, 0)
// }

// func TestGetPostsByUserID(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)
// 	post, err := addTestPost(user)
// 	posts := services.GetPostsByUserID(user.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, post, posts[0])
// }

// func TestGetPostsByUserIDEmpty(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)
// 	posts := services.GetPostsByUserID(user.ID)
// 	assert.NoError(t, err)
// 	assert.Empty(t, posts)
// }

// func TestFetchPost(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)
// 	post, err := addTestPost(user)
// 	fetchedPost, err := services.GetPostByID(post.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, post.ID, fetchedPost.ID)
// 	assert.Equal(t, post.Title, fetchedPost.Title)
// 	assert.Equal(t, post.Content, fetchedPost.Content)
// 	assert.Equal(t, user.ID, fetchedPost.UserID)
// }

// func TestFetchNotExistingPost(t *testing.T) {
// 	fetchedPost, err := services.GetPostByID(100)
// 	assert.Error(t, err)
// 	assert.Empty(t, fetchedPost)
// 	assert.Equal(t, "Post don't exists", err.Error())
// }

// func TestUpdatePost(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)

// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)

// 	post.Title = "New title"
// 	post.Content = "New content"
// 	_,err = services.UpdatePost(post.ID, post)
// 	assert.NoError(t, err)
// }

// func TestDeletePost(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)

// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)
// 	err = services.DeletePost(post.ID)
// 	assert.NoError(t, err)
// }
