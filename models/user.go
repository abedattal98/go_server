package models

import "time"

type User struct {
	ID         int
	Username   string `json:"username" binding:"required,min=5,max=30"`
	Password   string `json:"password" binding:"required,min=7,max=32"`
	Email      string `json:"email" validate:"required,email"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}
type UserDTO struct {
	ID         int    `json:"id,string,omitempty"`
	Username   string `json:"username" binding:"required,min=5,max=30"`
	Email      string `json:"email"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type CreateUserDTO struct {
	ID         int    `json:"id,string,omitempty"`
	Username   string `json:"username" binding:"required,min=5,max=30"`
	Password   string `json:"password" binding:"required,min=7,max=32"`
	Email      string `json:"email"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type UpdateUserDTO struct {
	Username   string `json:"username" binding:"required,min=5,max=30"`
	Password   string `json:"password" binding:"required,min=7,max=32"`
	Email      string `json:"email"`
	ModifiedAt time.Time
}

func ToUser(userDTO UserDTO) User {
	return User{Username: userDTO.Username, Email: userDTO.Email, CreatedAt: userDTO.CreatedAt, ModifiedAt: userDTO.ModifiedAt}
}

func ToUser2(userDTO CreateUserDTO) User {
	id := int(time.Now().Unix())
	return User{ID: id, Username: userDTO.Username, Email: userDTO.Email, CreatedAt: userDTO.CreatedAt, ModifiedAt: userDTO.ModifiedAt}
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{ID: user.ID, Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt, ModifiedAt: user.ModifiedAt}
}

func ToUsersDTOs(users []User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}
	return userdtos
}
