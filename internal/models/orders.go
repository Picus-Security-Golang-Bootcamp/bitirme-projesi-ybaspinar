package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null,foreignkey:UserID" json:"userId"`
	BasketID  uuid.UUID      `gorm:"type:uuid;not null,foreignkey:BasketID" json:"cartId""`
	Status    string         `json:"status"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()

	return
}
