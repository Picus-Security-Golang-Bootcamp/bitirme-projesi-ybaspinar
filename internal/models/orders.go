package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null,foreignkey:UserID"`
	BasketID  uuid.UUID      `gorm:"type:uuid;not null,foreignkey:UserID"`
	Status    string
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()

	return
}
