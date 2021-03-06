package middlewares

import (
	"errors"
	"net/http"
	"rgb/models"
	"rgb/services/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)


func Authorization(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing."})
		return
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is not valid."})
		return
	}
	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing bearer part."})
		return
	}
	userId, err := jwt.VerifyJWT(headerParts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// user, err := p.Repos.Users.FetchUser(userID)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }
	ctx.Set("userId", userId)
	ctx.Next()
}

func CurrentUser(ctx *gin.Context) (*models.User, error) {
	var err error
	_user, exists := ctx.Get("user")
	if !exists {
		err = errors.New("Current context user not set")
		return nil, err
	}
	user, ok := _user.(*models.User)
	if !ok {
		err = errors.New("Context user is not valid type")
		return nil, err
	}
	return user, nil
}
