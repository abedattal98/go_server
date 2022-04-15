package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rgb/models"
	"rgb/services"
	"rgb/services/jwt"
	"strconv"
)

func SignUp(ctx *gin.Context) {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if _,err := services.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": jwt.GenerateJWT(user),
	})
}

func SignIn(ctx *gin.Context) {
	user := new(models.User)
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

type UserAPI struct {
	UserService services.UserService
}

func ProvideUserAPI(p services.UserService) UserAPI {
	return UserAPI{UserService: p}
}

func (p *UserAPI) FindAll(c *gin.Context) {
	users := p.UserService.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": models.ToUsersDTOs(users)})
}

func (p *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(int(id))

	c.JSON(http.StatusOK, gin.H{"user": models.ToUserDTO(user)})
}

func (p *UserAPI) Create(c *gin.Context) {
	var createStudentDTO models.CreateUserDTO
	err := c.BindJSON(&createStudentDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdStudent := p.UserService.Save(models.ToUser2(createStudentDTO))
	c.JSON(http.StatusOK, gin.H{"student created successufuly": models.ToUserDTO(createdStudent)})
}

func (p *UserAPI) Update(c *gin.Context) {
	var userDTO models.UpdateUserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(int(id))
	if user == (models.User{}) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	user.Username = userDTO.Username
	user.Email = userDTO.Email
	updateStudent := p.UserService.Save(user)

	c.JSON(http.StatusOK, gin.H{"msg": "Post updated successfully.",
		"data":models.ToUserDTO(updateStudent)})
}

func (p *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(int(id))
	if user == (models.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.UserService.Delete(user)
	c.JSON(http.StatusOK, gin.H{"msg": "User deleted successfully."})
}
