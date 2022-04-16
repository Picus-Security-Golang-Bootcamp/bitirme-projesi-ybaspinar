package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"gorm.io/gorm"
)

type CategoriesRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoriesRepo {
	return &CategoriesRepo{db: db}
}

func (r *CategoriesRepo) Migrate() {
	r.db.AutoMigrate(&models.Category{})
}

func (r *CategoriesRepo) List() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoriesRepo) CreateFromCSV(categories []models.Category) error {
	err := r.db.Create(&categories).Error
	return err
}
