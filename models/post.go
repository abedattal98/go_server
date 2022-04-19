package models

import "time"

type Post struct {
	ID         int
	Title      string `binding:"required,min=3,max=50"`
	Content    string `binding:"required,min=5,max=5000"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	UserID     int `json:"userid" binding:"required"`
}
type UpdatePostDTO struct {
	Title      string 
	Content    string 
	ModifiedAt time.Time
}

func ToPostDTO(post Post) Post {
	return Post{ID: post.ID, Title: post.Title, Content: post.Content, UserID: post.UserID, CreatedAt: post.CreatedAt, ModifiedAt: post.ModifiedAt}
}
