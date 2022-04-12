package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"time"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null"`
	Basket      Basket
	Status      string
	UserAddress UserAddress
}

func (o *Order) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(o)
}
