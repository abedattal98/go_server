package server

import (
	"rgb/api"
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

