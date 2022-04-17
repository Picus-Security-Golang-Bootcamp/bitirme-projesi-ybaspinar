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
	zap.L().Info("Migrating products table")
	r.db.AutoMigrate(&models.Product{})
}

//Create creates a new product
func (r *ProductRepo) Create(product *models.Product) error {
	println(product)
	zap.L().Debug("ProductRepo.Create", zap.Any("product", product))
	return r.db.Create(product).Error
}

//GetAll returns all products
func (r *ProductRepo) GetAll(pageIndex, pageSize int) ([]models.Product, int) {
	zap.L().Debug("ProductRepo.GetAll")
	var products []models.Product
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

//Update updates a product
func (r *ProductRepo) Update(product *models.Product) error {
	zap.L().Debug("ProductRepo.Update", zap.Any("product", product))
	return r.db.Where(product.ID).Save(product).Error
}

//Delete deletes a product
func (r *ProductRepo) Delete(id string) error {
	zap.L().Debug("ProductRepo.Delete", zap.Any("id", id))
	return r.db.Delete(&models.Product{}, id).Error
}

//FuzzySearchSkuAndNameAndId returns products that match the given query
func (r *ProductRepo) FuzzySearchSkuAndNameAndId(search string, pageIndex, pageSize int) ([]models.Product, int) {
	zap.L().Debug("ProductRepo.FuzzySearchSkuAndNameAndId", zap.Any("search", search))
	var products []models.Product
	var count int64
	r.db.Offset((pageIndex-1)*pageSize).Limit(pageSize).Where("sku LIKE ? OR name LIKE ? OR id LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&products).Count(&count)
	return products, int(count)
}
