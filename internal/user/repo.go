package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
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

func (r *UserRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	return user, err
}
