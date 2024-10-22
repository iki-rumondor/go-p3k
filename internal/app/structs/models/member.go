package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID               uint              `gorm:"primaryKey"`
	Uuid             string            `gorm:"not_null;unique;size:64"`
	Name             string            `gorm:"not_null;size:128"`
	Position         string            `gorm:"not_null"`
	IsImportant      bool              `gorm:"not_null"`
	CreatedAt        int64             `gorm:"autoCreateTime:milli"`
	UpdatedAt        int64             `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	UserID           uint              `gorm:"not_null"`
	User             *User             `gorm:"constraint:OnDelete:CASCADE;"`
	MemberActivities *[]MemberActivity `gorm:"constraint:OnDelete:CASCADE;"`
}

func (m *Member) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
