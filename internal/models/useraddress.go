package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"time"
)

type UserAddress struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID    uuid.UUID `gorm:"type:uuid;not null;unique_index:idx_user_address_user_id_address_id"`
	AddressID uuid.UUID `gorm:"type:uuid;not null;unique_index:idx_user_address_user_id_address_id"`
}

func (u *UserAddress) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}
