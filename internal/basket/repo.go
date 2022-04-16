package basket

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"gorm.io/gorm"
)

type BasketRepo struct {
	db *gorm.DB
}

func NewBasketRepo(db *gorm.DB) *BasketRepo {
	return &BasketRepo{db}
}
func (r *BasketRepo) Migrate() {
	r.db.AutoMigrate(&models.Basket{})
}

func (r *BasketRepo) Create(basket *models.Basket) error {
	return r.db.Create(basket).Error
}

func (r *BasketRepo) GetByUserID(userID uint) (*models.Basket, error) {
	basket := &models.Basket{}
	err := r.db.Where("user_id = ?", userID).First(basket).Error
	return basket, err
}

func (r *BasketRepo) Update(basket *models.Basket) error {
	return r.db.Save(basket).Error
}
