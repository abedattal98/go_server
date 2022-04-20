package store

import (
	"rgb/models"
	"time"
)

var Posts = []models.Post {
	models.Post{
		ID:         1,
		Title:      "Gotham cronicles",
		Content:    "Joker is planning a big hit tonight.",
		CreatedAt:  time.Time{},
		ModifiedAt: time.Time{},
		UserID:     1,
	},
	models.Post{
		ID: 2,
		Title: "Justice league meeting",
		Content: "Darkseid is plotting again.",
		UserID: 1,
	},
}

