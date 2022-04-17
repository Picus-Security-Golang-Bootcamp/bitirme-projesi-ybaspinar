package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Migrate() {
	r.db.AutoMigrate(&models.User{})
}

//SignUp New user
func (r *UserRepo) SignUp(user *models.User) error {
	hPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hPassword)
	return r.db.Create(user).Error
}

//Login Checks given user credentials
func (r *UserRepo) Login(user *models.User) error {
	password := []byte(user.Password)
	r.db.Where("email = ? ", user.Email).First(user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		return err
	}
	return nil
}
