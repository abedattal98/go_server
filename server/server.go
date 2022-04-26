package server

import (
	"rgb/api"
	"rgb/repositories"
	"rgb/services"
	"rgb/services/jwt"
	// "rgb/conf"
)

func Start() {
	jwt.JwtSetup("test")

	repo := repositories.NewStorage(repositories.Memory)

	repos := repositories.NewRepositories(*repo)

	services := services.NewServices(services.Deps{
		Repos: &repos,
	})

	handlers := api.NewHandler(services)

	srv := handlers.Init()

	srv.Run()
}
