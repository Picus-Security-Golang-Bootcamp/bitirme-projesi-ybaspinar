package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoriesRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoriesRepo {
	return &CategoriesRepo{db: db}
}

func (r *CategoriesRepo) Migrate() {
	zap.L().Info("Migrating categories table")
	r.db.AutoMigrate(&models.Category{})
}

func (r *CategoriesRepo) GetAll(pageIndex, pageSize int) ([]models.Category, int) {
	zap.L().Debug("Getting all categories")
	var categories []models.Category
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)
	return categories, int(count)
}

func (r *CategoriesRepo) CreateFromCSV(categories *[]models.Category) error {
	zap.L().Debug("Creating categories from csv")
	err := r.db.Create(&categories).Error
	return err
}
