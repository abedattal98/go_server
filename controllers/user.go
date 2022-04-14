package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rgb/models"
	"rgb/services"
	"rgb/services/jwt"
)
// type AuthController struct {
// 	userService *services.UserService
// 	authService *services.AuthService
// }


func SignUp(ctx *gin.Context) {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	 userEntity,err := services.AddUser(user); 
	 if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": jwt.GenerateJWT(userEntity),
	})
}

func SignIn(ctx *gin.Context) {
	user := new(models.UserEntity)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	user, err := services.Authenticate(user.Username, user.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": jwt.GenerateJWT(user),
	})
}
