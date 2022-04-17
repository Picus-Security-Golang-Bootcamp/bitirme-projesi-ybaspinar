package orders

import (
	"errors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type OrdersRepo struct {
	db *gorm.DB
}

func NewOrdersRepo(db *gorm.DB) *OrdersRepo {
	return &OrdersRepo{db}
}

func (r *OrdersRepo) Migrate() {
	zap.L().Info("Migrating orders table")
	r.db.AutoMigrate(&models.Order{})
}

//Create Creates a new order
func (r *OrdersRepo) Create(order *models.Order) error {
	zap.L().Debug("Creating order", zap.Any("order", order))
	return r.db.Create(order).Error
}

//GetUserOrders gets all orders of a user
func (r *OrdersRepo) GetUsersOrders(userID string, pageIndex, pageSize int) ([]models.Order, int) {
	zap.L().Debug("GetUsersOrders", zap.String("userID", userID))
	var orders []models.Order
	var count int64
	r.db.Offset((pageIndex-1)*pageSize).Limit(pageSize).Where("user_id = ?", userID).Find(&orders).Count(&count)
	return orders, int(count)
}

//CancelOrder cancels an order if 14 days not passed
func (r *OrdersRepo) CancelOrder(orderID, userID string) error {
	zap.L().Debug("CancelOrder", zap.String("orderID", orderID))
	if r.CheckIf14DaysPassed(orderID) {
		return errors.New("14 days passed")
	}
	return r.db.Model(&models.Order{}).Where("id = ? AND userid = ?", orderID, userID).Update("status", "cancelled").Error
}

//CompleteOrder completes an order
func (r *OrdersRepo) CompleteOrder(orderID, userID string) error {
	zap.L().Debug("CompleteOrder", zap.String("orderID", orderID))
	return r.db.Model(&models.Order{}).Where("id = ? AND userid = ?", orderID, userID).Update("status", "delivered").Error
}

//CheckIf14DaysPassed checks if 14 days passed
func (r *OrdersRepo) CheckIf14DaysPassed(orderID string) bool {
	zap.L().Debug("CheckIf14DaysPassed", zap.String("orderID", orderID))
	var order models.Order
	r.db.Where("id = ?", orderID).First(&order)
	if order.CreatedAt.AddDate(0, 0, 14).Before(time.Now()) {
		return true
	}
	return false
}
