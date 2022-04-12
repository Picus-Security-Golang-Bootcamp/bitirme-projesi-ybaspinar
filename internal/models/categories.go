package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"time"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Products  []Product `gorm:"many2many:product_categories"`
}

func (c *Category) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(c)
}
