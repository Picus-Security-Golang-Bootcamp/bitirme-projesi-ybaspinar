package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Basket struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null,foreignkey:UserID" json:"userid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Products  []Product      `gorm:"many2many:basket_products" json:"products"`
}

func (b *Basket) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()

	return
}
