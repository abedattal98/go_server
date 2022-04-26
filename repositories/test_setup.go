package repositories

import (
	"rgb/models"
	"time"
)

func addTestUser(repo Repositories) (models.User, error) {

	user := models.User{
		Username: "rand.String(10)",
		Password: "secret123",
		ID:       int(time.Now().Unix()),
	}
	user, err := repo.Users.Save(user)
	return user, err
}

func addTestPost(user models.User, repo Repositories) (models.Post, error) {
	post := models.Post{
		Title:   "Gotham cronicles",
		Content: "Joker is planning big hit tonight.",
	}
	addPost, err := repo.Posts.AddPost(user.ID, post)
	return addPost, err
}

func NewRepositoriesTest(t Type) Repositories {
	repo := NewStorage(t)
	repos := NewRepositories(*repo)

	return repos
}