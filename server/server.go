package server

import (
	"rgb/controllers"
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

	handlers := controllers.NewHandler(services)

	srv := handlers.Init()

	srv.Run()
}
