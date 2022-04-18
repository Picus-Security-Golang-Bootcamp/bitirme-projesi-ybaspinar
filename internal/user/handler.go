package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	jwtHelper "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/JWT"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

type UserHandler struct {
	repo *UserRepo
	cfg  *config.Config
}

//SignUp Creates new user with given data and returns token
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
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":  user.ID,
		"email":   user.Email,
		"iat":     time.Now().Unix(),
		"iss":     os.Getenv("ENV"),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"isAdmin": user.IsAdmin,
	})
	token := jwtHelper.GenerateToken(jwtClaims, h.cfg.JWTConfig.SecretKey)
	context.JSON(http.StatusOK, token)
}

//Login compares user's email and password with database and returns token
func (h UserHandler) Login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Login(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":  user.ID,
		"email":   user.Email,
		"iat":     time.Now().Unix(),
		"iss":     os.Getenv("ENV"),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"isAdmin": user.IsAdmin,
	})
	token := jwtHelper.GenerateToken(jwtClaims, h.cfg.JWTConfig.SecretKey)
	context.JSON(http.StatusOK, token)
}

func NewUserHandler(r *gin.RouterGroup, repo *UserRepo, cfg *config.Config) {
	handler := &UserHandler{repo: repo, cfg: cfg}
	r.POST("/signup", handler.SignUp)
	r.POST("/login", handler.Login)
}
