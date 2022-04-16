package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserAddress struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Address   string
	UserID    User `gorm:"foreignkey:UserID"`
}

func (u *UserAddress) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
