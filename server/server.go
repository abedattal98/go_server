package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rgb/controllers"
	"rgb/services"
	"rgb/middlewares"
	// "rgb/conf"

)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	// Create API route group
	api := router.Group("/api")
	{
		api.POST("/signup", controllers.SignUp)
		api.POST("/signin", controllers.SignIn)



		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})
	}
	authorized := api.Group("/")
	authorized.Use(middlewares.Authorization)
	{
		authorized.POST("/posts", controllers.CreatePost)
		authorized.GET("/posts", controllers.GetPostsByUserID)
		authorized.GET("/posts/:id", controllers.GetPostById)
		authorized.DELETE("/posts/:id", controllers.DeletePost)
		authorized.PUT("/posts/:id", controllers.UpdatePost)

	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func Start() {
	// services.JwtSetup(conf.NewConfig().JwtSecret)
	services.JwtSetup("test")

	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
