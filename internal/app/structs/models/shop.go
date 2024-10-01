package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shop struct {
	ID            uint   `gorm:"primaryKey"`
	Uuid          string `gorm:"not_null;unique;size:64"`
	Name          string `gorm:"not_null;size:128"`
	Owner         string `gorm:"not_null"`
	Address       string `gorm:"not_null"`
	PhoneNumber   string `gorm:"not_null"`
	ShopImage     string `gorm:"not_null"`
	IdentityImage string `gorm:"not_null"`
	CreatedAt     int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt     int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	UserID        uint   `gorm:"not_null"`
	User          *User
}

func (m *Shop) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
