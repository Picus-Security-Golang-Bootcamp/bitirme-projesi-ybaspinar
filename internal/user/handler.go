package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	repo *UserRepo
}

//TODO: implement JWT
func (h UserHandler) SignUp(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.SignUp(&user); err != nil {
		context.JSON(400, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

//TODO: implement JWT
func (h UserHandler) Login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Login(&user); err != nil {
		context.JSON(400, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}

func NewUserHandler(r *gin.RouterGroup, repo *UserRepo) {
	handler := &UserHandler{
		repo: repo,
	}
	r.POST("/signup", handler.SignUp)
	r.POST("/login", handler.Login)
}
