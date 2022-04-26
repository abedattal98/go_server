package server

import (
	"rgb/api"
	"rgb/models"
	"rgb/repositories"
	"rgb/services"
	"rgb/services/jwt"
	// "rgb/conf"
)

func Start() {
	//setup jwt client
	jwt.JwtSetup("test")

	// create new memory storage
	repo := repositories.NewStorage(repositories.Memory)

	//initialize  repository
	repos := repositories.NewRepositories(*repo)

	//Populate default users
	PopulateUsers(repos)

	//initialize  services
	services := services.NewServices(services.Deps{
		Repos: &repos,
	})
	//initialize controllers with services
	handlers := api.NewHandler(services)

	//initialize routes
	srv := handlers.Init()

	srv.Run()
}

// PopulateUsers populates the Users variable with User
func PopulateUsers(repo repositories.Repositories) {
	defaultUser := models.User{
        ID:       1,
        Email:    "admin@admin.com",
        Password: "admin",
		Username: "admin",
    }
	repo.Users.Save(defaultUser)
}
