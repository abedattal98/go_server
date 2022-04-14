package models

type Post struct {
  Title      string    `json:"title" binding:"required,min=3,max=50"`
  Content    string    `json:"content" binding:"required,min=5,max=5000"`
  UserID     int       `json:"userId" binding:"required"`
}
type PostEntity struct {
	Entity
	Post
}
