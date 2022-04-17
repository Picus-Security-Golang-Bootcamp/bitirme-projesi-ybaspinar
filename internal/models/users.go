package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FirstName string         `gorm:"type:varchar(100);not null" json:"firstname"`
	LastName  string         `gorm:"type:varchar(100);not null" json:"lastname"`
	Password  string         `gorm:"type:varchar(100);not null" json:"password"`
	Email     string         `gorm:"type:varchar(100);not null;unique" json:"email"`
	IsAdmin   bool           `gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
