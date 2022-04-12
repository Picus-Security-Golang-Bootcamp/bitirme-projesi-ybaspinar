package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string
	SKU         string
	Description string
	Price       float64
	Stock       int
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}
