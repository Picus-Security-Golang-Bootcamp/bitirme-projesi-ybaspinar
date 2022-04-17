package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint `gorm:"primaryKey" json:"id" csv:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(100);unique_index" json:"name" csv:"name"`
}
