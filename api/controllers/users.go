package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"rgb/models"
	"rgb/services/jwt"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	api.POST("/signup", h.SignUp)
	api.POST("/signin", h.SignIn)
	users := api.Group("/users")
	{
		users.GET("/", h.FindAll)
		users.GET("/:id", h.FindById)
		users.POST("/", h.Create)
		users.PUT("/:id", h.Update)
		users.DELETE("/:id", h.Delete)
	}
}

func (h *Handler) FindAll(c *gin.Context) {
	users, error := h.services.Users.FindAll()
	if error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": models.ToUsersDTOs(users)})
}

func (h *Handler) FindById(c *gin.Context) {
	//get user Id from param
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.services.Users.FindById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": models.ToUserDTO(user)})
}

func (h *Handler) Create(c *gin.Context) {
	//bind the user to the request body
	var createStudentDTO models.CreateUserDTO
	err := c.BindJSON(&createStudentDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createStudentDTO.CreatedAt = time.Now().UTC()
	createStudentDTO.ModifiedAt = time.Now().UTC()

	//save user data
	createdStudent, err := h.services.Users.Save(models.ToUser2(createStudentDTO))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "user created successfully", "user": models.ToUserDTO(createdStudent)})
}

func (h *Handler) Update(c *gin.Context) {
	var userDTO models.UpdateUserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.services.Users.FindById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = userDTO.Username
	user.Email = userDTO.Email
	user.ModifiedAt = time.Now().UTC()

	updateStudent, err := h.services.Users.Save(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "user updated successfully.",
		"data": models.ToUserDTO(updateStudent)})
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.services.Users.FindById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == (models.User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	h.services.Users.Delete(user)
	c.JSON(http.StatusOK, gin.H{"msg": "user deleted successfully."})
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var createStudentDTO models.CreateUserDTO
	//assign values from body to the createStudentDTO
	err := ctx.BindJSON(&createStudentDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createStudentDTO.CreatedAt = time.Now().UTC()
	createStudentDTO.ModifiedAt = time.Now().UTC()

	//send values to the service
	createdUser, err := h.services.Users.Save(models.ToUser2(createStudentDTO))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//return the created user jwt
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": jwt.GenerateJWT(createdUser),
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var LoginDTO models.LoginDTO
	//assign values from body to the LoginDTO
	err := ctx.BindJSON(&LoginDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//validate the LoginDTO
	loginUser, err := h.services.Users.Authenticate(LoginDTO.Email, LoginDTO.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
		return
	}
	//return the created user jwt
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Signed in successfully.",
		"jwt":  jwt.GenerateJWT(loginUser),
		"user": loginUser,
	})
}
