package controllers

import (
	"net/http"
	"rgb/middlewares"
	"rgb/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initPostsRoutes(api *gin.RouterGroup) {
	users := api.Group("/users/:id/posts")
	{
		users.GET("/", h.GetPostsByUserID)
		users.POST("/", h.CreatePost)
	}
	posts := api.Group("/posts")
	{
		posts.GET("/:id", h.GetPostById)
		posts.PUT("/:id", h.Update)
		posts.DELETE("/:id", h.Delete)
	}
	authorized := api.Use(middlewares.Authorization)
	// Init router
	authorized.GET("/pinggg", func(c *gin.Context) {
		c.String(http.StatusOK, "ponggg")
	})

}

func (h *Handler) CreatePost(ctx *gin.Context) {
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

	//send the post data to the save h.services.Posts
	createdPost, err := h.services.Posts.AddPost(userId, *post)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post created successfully.",
		"data": createdPost,
	})
}

func (h *Handler) UpdatePost(ctx *gin.Context) {
	var updatePostDTO models.UpdatePostDTO
	err := ctx.BindJSON(&updatePostDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, err := h.services.Posts.GetPostByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	post.Title = updatePostDTO.Title
	post.Content = updatePostDTO.Content
	post.ModifiedAt = time.Now().UTC()

	updatedPost, err := h.services.Posts.UpdatePost(id, post)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Post updated successfully.",
		"data": models.ToPostDTO(updatedPost)})
}

//get posts by user ID
func (h *Handler) GetPostsByUserID(ctx *gin.Context) {
	//get user id
	userParamId := ctx.Param("id")
	userId, err := strconv.Atoi(userParamId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	//check all posts to the related user id
	posts := h.services.Posts.GetPostsByUserID(userId)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Posts fetched successfully.",
		"data": posts,
	})
}

func (h *Handler) GetPostById(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	//check if post exists
	post, err := h.services.Posts.GetPostByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post fetched successfully.",
		"data": post,
	})
}

func (h *Handler) DeletePost(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not valid ID."})
		return
	}
	//check if post exists

	post, err := h.services.Posts.GetPostByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	//delete if post exists

	if err := h.services.Posts.DeletePost(post.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Post deleted successfully."})
}
