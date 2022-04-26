package server

import (
	"net/http"
	"rgb/controllers"
	"rgb/services"

	"rgb/repositories"
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
	srv := NewServer(handlers.Init())
	srv.Run()
}

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
