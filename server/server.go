package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rgb/controllers"
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

    api.POST("/posts", controllers.CreatePost)
		api.GET("/posts", controllers.GetPosts)

		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func Start() {
	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
