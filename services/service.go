package services

import (
	"rgb/interfaces"
	"rgb/repositories"
)

type Services struct {
	Users interfaces.UserService
	Posts interfaces.PostService
}

type Deps struct {
	Repos *repositories.Repositories
}

func NewServices(deps Deps) *Services {

	usersService := NewUsersService(deps.Repos.Users)
	postsService := NewPostsService(deps.Repos.Posts)

	return &Services{
		Users: usersService,
		Posts: postsService,
	}
}
