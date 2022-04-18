package controllers

import (
	"net/http"
	"rgb/models"
	interfaceUser "rgb/repositories/interface"
	"rgb/services"
	"rgb/services/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if _, err := services.AddUser(user); err != nil {
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
	UserService interfaceUser.IUserService
}

func ProvideUserAPI(usecase interfaceUser.IUserService) *UserAPI {
	return &UserAPI{
		UserService: usecase,
	}
}

func (p *UserAPI) FindAll(c *gin.Context) {
	users, error := p.UserService.FindAll()
	if error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": models.ToUsersDTOs(users)})
}

func (p *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := p.UserService.FindByID(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": models.ToUserDTO(user)})
}

func (p *UserAPI) Create(c *gin.Context) {
	var createStudentDTO models.CreateUserDTO
	err := c.BindJSON(&createStudentDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdStudent, err := p.UserService.Save(models.ToUser2(createStudentDTO))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	user, err := p.UserService.FindByID(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = userDTO.Username
	user.Email = userDTO.Email
	user.ModifiedAt = userDTO.ModifiedAt

	updateStudent, err := p.UserService.Save(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Post updated successfully.",
		"data": models.ToUserDTO(updateStudent)})
}

func (p *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := p.UserService.FindByID(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == (models.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.UserService.Delete(user)
	c.JSON(http.StatusOK, gin.H{"msg": "User deleted successfully."})
}
