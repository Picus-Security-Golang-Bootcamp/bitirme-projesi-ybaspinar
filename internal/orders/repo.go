package orders

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"gorm.io/gorm"
)

type OrdersRepo struct {
	db *gorm.DB
}

func NewOrdersRepo(db *gorm.DB) *OrdersRepo {
	return &OrdersRepo{db}
}

func (r *OrdersRepo) Migrate() {
	r.db.AutoMigrate(&models.Order{})
}

func (r *OrdersRepo) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrdersRepo) GetUsersOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrdersRepo) CancelOrder(orderID uint) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", "cancelled").Error
}
