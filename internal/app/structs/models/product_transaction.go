package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductTransaction struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null;unique;size:64"`
	Quantity   int64  `gorm:"not_null"`
	IsResponse bool   `gorm:"not_null"`
	IsAccept   bool   `gorm:"not_null"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	UserID     uint   `gorm:"not_null"`
	ProductID  uint   `gorm:"not_null"`
	Product    *Product
	User       *User
}

func (m *ProductTransaction) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
