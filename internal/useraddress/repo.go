package useraddress

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"gorm.io/gorm"
)

type UserAddressRepo struct {
	db *gorm.DB
}

func NewUserAddressRepo(db *gorm.DB) *UserAddressRepo {
	return &UserAddressRepo{db: db}
}

func (r *UserAddressRepo) Migrate() {
	r.db.AutoMigrate(&models.UserAddress{})
}

func (r *UserAddressRepo) Create(userAddress *models.UserAddress) error {
	return r.db.Create(userAddress).Error
}

func (r *UserAddressRepo) FindByUserID(userID uint) ([]models.UserAddress, error) {
	var userAddresses []models.UserAddress
	err := r.db.Where("user_id = ?", userID).Find(&userAddresses).Error
	return userAddresses, err
}
func (r *UserAddressRepo) FindByID(id uint) (*models.UserAddress, error) {
	var userAddress models.UserAddress
	err := r.db.Where("id = ?", id).First(&userAddress).Error
	return &userAddress, err
}
func (r *UserAddressRepo) Update(userAddress *models.UserAddress) error {
	return r.db.Save(userAddress).Error
}
func (r *UserAddressRepo) Delete(userAddress *models.UserAddress) error {
	return r.db.Delete(userAddress).Error
}
