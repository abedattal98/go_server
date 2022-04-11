package models
import "time"

type User struct {
  ID             int
  Username string `binding:"required,min=5,max=30"`
  Password string `binding:"required,min=7,max=32"`
  CreatedAt      time.Time
	ModifiedAt     time.Time
  Posts          []*Post `json:"-"`
}
