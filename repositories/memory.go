package repositories

import (
	"rgb/models"
)

// Memory data storage layered save only in memory
type MemoryStorage struct {
	users  []models.User
	posts []models.Post
}
