package controllers

import (
	"net/http"
	"rgb/models"
	"rgb/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	//get user id from param
	userParamId := ctx.Param("id")
	//convert user id to integer
	userId, err := strconv.Atoi(userParamId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	//create new post from model
	post := new(models.Post)
	post.UserID = userId
	if err := ctx.Bind(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	post.CreatedAt = time.Now().UTC()
	post.ModifiedAt = time.Now().UTC()

	//send the post data to the save services
	if err := services.AddPost(userId, post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post created successfully.",
		"data": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	var updatePostDTO models.UpdatePostDTO
	err := ctx.BindJSON(&updatePostDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, err := services.GetPostByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	post.Title = updatePostDTO.Title
	post.Content = updatePostDTO.Content
	post.ModifiedAt = time.Now().UTC()

	if err := services.UpdatePost(id, post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Post updated successfully.",
		"data": models.ToPostDTO(*post)})
}

//get posts by user ID
func GetPostsByUserID(ctx *gin.Context) {
	//get user id
	userParamId := ctx.Param("id")
	userId, err := strconv.Atoi(userParamId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	//check all posts to the related user id
	posts := services.GetPostsByUserID(userId)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Posts fetched successfully.",
		"data": posts,
	})
}

func GetPostById(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	//check if post exists
	post, err := services.GetPostByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post fetched successfully.",
		"data": post,
	})
}

func DeletePost(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	//check if post exists

	post, err := services.GetPostByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	//delete if post exists

	if err := services.DeletePost(post.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Post deleted successfully."})
}
