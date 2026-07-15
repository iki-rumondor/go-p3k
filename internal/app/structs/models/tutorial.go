package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tutorial struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique;size:64"`
	Title     string `gorm:"not_null;size:255"`
	Content   string `gorm:"type:text;not_null"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}

func (m *Tutorial) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
