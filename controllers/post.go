package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rgb/models"
	"rgb/store"
)

func CreatePost(ctx *gin.Context) {
	post := new(models.Post)
	if err := ctx.Bind(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	store.Posts = append(store.Posts, post)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post created successfully.",
	})
}

func GetPost(ctx *gin.Context) {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	for _, u := range store.Users {
		if u.Username == user.Username && u.Password == user.Password {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Signed in successfully.",
				"jwt": "123456789",
			})
			return
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed."})
}

func GetPosts(ctx *gin.Context) {
	posts := store.Posts;
	ctx.JSON(http.StatusOK, 
		posts , 
	)
	return
}
