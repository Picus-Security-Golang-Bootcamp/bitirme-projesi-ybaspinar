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

// Create creates a new basket
func (r *BasketRepo) Create(basket *models.Basket) error {
	return r.db.Create(basket).Error
}

// GetAllByUserID returns all baskets of a user
func (r *BasketRepo) GetAllByUserID(userID string, pageIndex, pageSize int) (models.Basket, int) {
	basket := models.Basket{}
	var count int64
	r.db.Preload("basket_products").Offset((pageIndex-1)*pageSize).Limit(pageSize).Where("user_id = ?", userID).Find(&basket).Count(&count)
	return basket, int(count)
}

// update updates a basket
func (r *BasketRepo) Update(basket *models.Basket) error {
	return r.db.Where(basket.ID).Save(basket).Error
}

// delete deletes a basket
func (r *BasketRepo) Delete(basket *models.Basket) error {
	return r.db.Where("id = ?", basket.ID).Delete(basket).Error
}
