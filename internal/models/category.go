package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Products  []Product `gorm:"many2many:product_categories"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()

	return
}
