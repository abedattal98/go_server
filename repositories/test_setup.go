package repositories

import (
	"rgb/domain"
	"rgb/models"
	 "rgb/services"
	"time"
)
func AddTestUser(repo domain.IUserRepository) (models.User, error) {
	user := models.User{
		Username: "rand.String(10)",
		Password: "secret123",
		ID:       int(time.Now().Unix()),
	}
	user, err := repo.Save(user)
	return user, err
}
func addTestPost(user models.User) (models.Post, error) {
	post := models.Post{
		Title:   "Gotham cronicles",
		Content: "Joker is planning big hit tonight.",
	}
	addPost,err := services.AddPost(user.ID, post)
	return addPost, err
}