package services

import (
	"rgb/models"
	"rgb/store"

	"github.com/gin-gonic/gin"
)

func testSetup() {
	gin.SetMode(gin.TestMode)
	store.Users=nil;
}

func addTestUser() (*models.User, error) {
	user := &models.User{
		Username: "rand.String(10)",
		Password: "secret123",
	}
	user,err := AddUser(user)
	return user, err
}

func addTestPost(user *models.User) (*models.Post, error) {
	post := &models.Post{
		Title:   "Gotham cronicles",
		Content: "Joker is planning big hit tonight.",
	}
	err := AddPost(user, post)
	return post, err
}
