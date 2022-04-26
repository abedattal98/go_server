package repositories

import "rgb/interfaces"

type Repositories struct {
	Users interfaces.UserRepository
	Posts interfaces.PostRepository
}

func NewRepositories(db MemoryStorage) Repositories {
	return Repositories{
		Users: NewUsersRepo(db),
		Posts: NewPostsRepo(db),
	}
}
