package products

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db}
}
func (r *ProductRepo) Migrate() {
	r.db.AutoMigrate(&models.Product{})
}

func (r *ProductRepo) Create(product *models.Product) error {
	zap.L().Debug("ProductRepo.Create", zap.Any("product", product))
	return r.db.Create(product).Error
}

func (r *ProductRepo) GetAll() ([]models.Product, error) {
	zap.L().Debug("ProductRepo.GetAll")
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepo) GetByID(id string) (*models.Product, error) {
	zap.L().Debug("ProductRepo.GetByID", zap.Any("id", id))
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepo) Update(product *models.Product) error {
	zap.L().Debug("ProductRepo.Update", zap.Any("product", product))
	return r.db.Save(product).Error
}

func (r *ProductRepo) Delete(id string) error {
	zap.L().Debug("ProductRepo.Delete", zap.Any("id", id))
	return r.db.Delete(&models.Product{}, id).Error
}
