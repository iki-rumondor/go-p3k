package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Uuid        string `gorm:"not_null;unique;size:64"`
	Name        string `gorm:"not_null;size:128"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}

func (m *Category) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
