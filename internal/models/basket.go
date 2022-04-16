package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Basket struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Products    []Product      `gorm:"many2many:basket_products"`
	TotalAmount float64        `gorm:"type:decimal(10,2);not null"`
}

func (b *Basket) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()

	return
}
