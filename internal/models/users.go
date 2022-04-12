package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"time"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName     string
	LastName      string
	Password      string
	Email         string
	UserAddresses []UserAddress
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (u *User) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}
