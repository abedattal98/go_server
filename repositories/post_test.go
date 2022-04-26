package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPostsByUserID(t *testing.T) {
	repos:= NewRepositoriesTest(Memory)

	user, err := addTestUser(repos)
	assert.NoError(t, err)
	post, err := addTestPost(user, repos)
	posts := repos.Posts.GetPostsByUserID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, post, posts[0])
}

func TestGetPostsByUserIDEmpty(t *testing.T) {
	repos:= NewRepositoriesTest(Memory)

	user, err := addTestUser(repos)
	assert.NoError(t, err)
	posts := repos.Posts.GetPostsByUserID(user.ID)
	assert.NoError(t, err)
	assert.Empty(t, posts)
}

func TestFetchPost(t *testing.T) {
	repos:= NewRepositoriesTest(Memory)

	user, err := addTestUser(repos)
	assert.NoError(t, err)
	post, err := addTestPost(user, repos)
	fetchedPost, err := repos.Posts.GetPostByID(post.ID)
	assert.NoError(t, err)
	assert.Equal(t, post.ID, fetchedPost.ID)
	assert.Equal(t, post.Title, fetchedPost.Title)
	assert.Equal(t, post.Content, fetchedPost.Content)
	assert.Equal(t, user.ID, fetchedPost.UserID)
}

func TestFetchNotExistingPost(t *testing.T) {
	repos:= NewRepositoriesTest(Memory)


	fetchedPost, err := repos.Posts.GetPostByID(100)
	assert.Error(t, err)
	assert.Empty(t, fetchedPost)
	assert.Equal(t, "Post don't exists", err.Error())
}

func TestUpdatePost(t *testing.T) {
	repos:= NewRepositoriesTest(Memory)

	user, err := addTestUser(repos)
	assert.NoError(t, err)
	post, err := addTestPost(user, repos)
	assert.NoError(t, err)
	post.Title = "New title"
	post.Content = "New content"
	_, err = repos.Posts.UpdatePost(post.ID, post)
	assert.NoError(t, err)
}

func TestDeletePost(t *testing.T) {
	repos:= NewRepositoriesTest(Memory)

	user, err := addTestUser(repos)
	assert.NoError(t, err)
	post, err := addTestPost(user, repos)
	assert.NoError(t, err)
	err = repos.Posts.DeletePost(post.ID)
	assert.NoError(t, err)
}
