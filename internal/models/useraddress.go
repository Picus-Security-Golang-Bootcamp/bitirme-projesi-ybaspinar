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

	UserID    uuid.UUID `gorm:"type:uuid;not null;unique_index:idx_user_address_user_id_address_id"`
	AddressID uuid.UUID `gorm:"type:uuid;not null;unique_index:idx_user_address_user_id_address_id"`
}
