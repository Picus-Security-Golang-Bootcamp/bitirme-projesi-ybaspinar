package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	FirstName     string
	LastName      string
	Password      string
	Email         string
	UserAddresses []UserAddress
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	IsAdmin       bool           `gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
