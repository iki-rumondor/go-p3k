package models

import (
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberActivity struct {
	ID              uint   `gorm:"primaryKey"`
	Uuid            string `gorm:"not_null;unique;size:64"`
	MemberID        uint   `gorm:"not_null"`
	ActivityID      uint   `gorm:"not_null"`
	AttendenceImage string `gorm:"not_null"`
	IsAccept        bool   `gorm:"not_null"`
	CreatedAt       int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt       int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Activity        *Activity
	Member          *Member
}

func (m *MemberActivity) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *MemberActivity) BeforeDelete(tx *gorm.DB) error {
	folder := "internal/files/attendances"
	pathFile := filepath.Join(folder, m.AttendenceImage)
	if err := os.Remove(pathFile); err != nil {
		log.Println(err.Error())
	}
	return nil
}
