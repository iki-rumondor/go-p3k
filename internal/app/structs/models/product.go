package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique;size:64"`
	Name      string `gorm:"not_null;size:128"`
	Price     int64  `gorm:"not_null"`
	Stock     int64  `gorm:"not_null"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	ShopID    uint   `gorm:"not_null"`
	Shop      *Shop
}

func (m *Product) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
