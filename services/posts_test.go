package services

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestAddPost(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)

// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)
// 	assert.Greater(t, post.ID, 0)
// }

// func TestGetPostsByUserID(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)
// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)
// 	posts := GetPostsByUserID(user)
// 	assert.NoError(t, err)
// 	assert.Equal(t, post, posts[0])
// }

// func TestGetPostsByUserIDEmpty(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)

// 	Posts := GetPostsByUserID(user)
// 	assert.NoError(t, err)
// 	assert.Empty(t, Posts)
// }

// func TestFetchPost(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)
// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)

// 	fetchedPost, err := GetPostByID(post.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, post.ID, fetchedPost.ID)
// 	assert.Equal(t, post.Title, fetchedPost.Title)
// 	assert.Equal(t, post.Content, fetchedPost.Content)
// 	assert.Equal(t, user.ID, fetchedPost.UserID)
// }

// func TestFetchNotExistingPost(t *testing.T) {
// 	testSetup()

// 	fetchedPost, err := GetPostByID(1)
// 	assert.Error(t, err)
// 	assert.Nil(t, fetchedPost)
// 	assert.Equal(t, "Post don't exists", err.Error())
// }

// func TestUpdatePost(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)
// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)

// 	post.Title = "New title"
// 	post.Content = "New content"
// 	err = UpdatePost(post.ID,post)
// 	assert.NoError(t, err)
// }

// func TestDeletePost(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)
// 	post, err := addTestPost(user)
// 	assert.NoError(t, err)

// 	err = DeletePost(post.ID)
// 	assert.NoError(t, err)
// }
