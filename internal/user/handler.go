package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *UserRepo
}

//TODO: implement
func (h UserHandler) SignUp(context *gin.Context) {

}

//TODO: implement
func (h UserHandler) Login(context *gin.Context) {

}

func NewUserHandler(r *gin.RouterGroup, repo *UserRepo) {
	handler := &UserHandler{
		repo: repo,
	}
	r.POST("/signup", handler.SignUp)
	r.POST("/login", handler.Login)
}
