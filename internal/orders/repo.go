package orders

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"go.uber.org/zap"
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

func (r *OrdersRepo) GetUsersOrders(userID string, pageIndex, pageSize int) ([]models.Order, int) {
	zap.L().Debug("GetUsersOrders", zap.String("userID", userID))
	var orders []models.Order
	var count int64
	r.db.Offset((pageIndex-1)*pageSize).Limit(pageSize).Where("user_id = ?", userID).Find(&orders).Count(&count)
	return orders, int(count)
}

func (r *OrdersRepo) CancelOrder(orderID string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", "cancelled").Error
}

func (r *OrdersRepo) ConfirmOrder(orderID string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", "confirmed").Error
}

func (r *OrdersRepo) CompleteOrder(orderID string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", "delivered").Error
}
