package repositories

import (
	"rgb/models"
)

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users: []models.User{
			{
				ID:       1,
				Email:    "admin@admin.com",
				Password: "admin",
				Username: "admin",
			},
		},
		posts: []models.Post{
			{
				ID:      1,
				Title:   "Gotham cronicles",
				Content: "Joker is planning big hit tonight.",
				UserID:  1,
			},
		},
	}
}

// Memory data storage layered save only in memory
type MemoryStorage struct {
	users []models.User
	posts []models.Post
}
