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

func (r *BasketRepo) GetAllByUserID(userID string, pageIndex, pageSize int) ([]models.Basket, int) {
	basket := []models.Basket{}
	var count int64
	r.db.Offset((pageIndex-1)*pageSize).Limit(pageSize).Where("user_id = ?", userID).Find(&basket).Count(&count)
	return basket, int(count)
}

func (r *BasketRepo) Update(basket *models.Basket) error {
	return r.db.Where(basket.ID).Save(basket).Error
}

func (r *BasketRepo) Delete(basket *models.Basket) error {
	return r.db.Delete(basket).Error
}
