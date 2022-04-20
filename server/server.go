package server

import (
	"net/http"
	"rgb/controllers"

	"rgb/middlewares"
	"rgb/repositories"
	"rgb/services"
	"rgb/services/jwt"

	"github.com/gin-gonic/gin"
	// "rgb/conf"
)

var UserRepository = repositories.ProvideUserRepository()

func initUserAPI() controllers.UserAPI {
	studentService := services.ProvideUserService(UserRepository)
	userAPI := controllers.ProvideUserAPI(studentService)
	return *userAPI
}
func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true
	userAPI := initUserAPI()

	// Create API route group
	api := router.Group("/api")
	{
		api.POST("/signup", userAPI.SignUp)
		api.POST("/signin", userAPI.SignIn)

		api.GET("/users", userAPI.FindAll)
		api.GET("/users/:id", userAPI.FindByID)
		api.POST("/users", userAPI.Create)
		api.PUT("/users/:id", userAPI.Update)
		api.DELETE("/users/:id", userAPI.Delete)



		api.GET("/users/:id/posts", controllers.GetPostsByUserID)
		api.POST("/users/:id/posts", controllers.CreatePost)
		api.GET("/posts/:id", controllers.GetPostById)
		api.DELETE("/posts/:id", controllers.DeletePost)
		api.PUT("/posts/:id", controllers.UpdatePost)
	}
	authorized := api.Group("/")
	authMiddleware := middlewares.AuthMiddleware{Repo:UserRepository}
	authorized.Use(authMiddleware.Authorization)
	authorized.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "world"})
	})
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func Start() {
	// services.JwtSetup(conf.NewConfig().JwtSecret)
	jwt.JwtSetup("test")

	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
