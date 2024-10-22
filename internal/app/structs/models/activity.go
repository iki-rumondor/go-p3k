package models

import (
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	ID            uint              `gorm:"primaryKey"`
	Uuid          string            `gorm:"not_null;unique;size:64"`
	Group         uint              `gorm:"not_null"`
	Title         string            `gorm:"not_null"`
	Description   string            `gorm:"not_null"`
	ImageName     string            `gorm:"not_null;"`
	CreatedUserID uint              `gorm:"not_null"`
	UpdatedUserID uint              `gorm:"not_null"`
	CreatedAt     int64             `gorm:"autoCreateTime:milli"`
	UpdatedAt     int64             `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	CreatedUser   *User             `gorm:"foreignKey:CreatedUserID"`
	UpdatedUser   *User             `gorm:"foreignKey:UpdatedUserID"`
	Members       *[]MemberActivity `gorm:"constraint:OnDelete:CASCADE;"`
}

func (m *Activity) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *Activity) BeforeDelete(tx *gorm.DB) error {
	folder := "internal/files/activities"
	pathFile := filepath.Join(folder, m.ImageName)
	if err := os.Remove(pathFile); err != nil {
		log.Println(err.Error())
	}
	return nil
}
