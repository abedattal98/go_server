package repositories

import "rgb/interfaces"

type Repositories struct {
	Users interfaces.IUserRepository
	Posts interfaces.IPostRepository
}

func NewRepositories(db MemoryStorage) Repositories {
	return Repositories{
		Users: NewUsersRepo(db),
		Posts: NewPostsRepo(db),
	}
}
