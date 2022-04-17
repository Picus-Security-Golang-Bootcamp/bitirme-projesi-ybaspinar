package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name"`
	SKU         string         `json:"sku"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`
	CategoryID  int            `gomr:"fk:category" json:"category_id"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()

	return
}

// json to struct binding
