package models

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
	"github.com/iki-rumondor/go-p3k/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID                  uint                  `gorm:"primaryKey"`
	Uuid                string                `gorm:"not_null;unique;size:64"`
	Name                string                `gorm:"not_null;size:128"`
	Username            string                `gorm:"not_null;unique;size:32"`
	Password            string                `gorm:"not_null;size:64"`
	Active              bool                  `gorm:"not_null"`
	RoleID              uint                  `gorm:"not_null"`
	CreatedAt           int64                 `gorm:"autoCreateTime:milli"`
	UpdatedAt           int64                 `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	ProductTransactions *[]ProductTransaction `gorm:"constraint:OnDelete:CASCADE;"`
	Role                *Role
	Shop                *Shop
	Guest               *Guest
	Member              *Member
	Citizen             *Citizen
}

func (m *User) BeforeSave(tx *gorm.DB) error {
	if result := tx.First(&User{}, "username = ? AND id != ?", m.Username, m.ID).RowsAffected; result > 0 {
		return response.BADREQ_ERR("Username Yang Dipakai Sudah Terdaftar Pada Akun Lain")
	}

	return nil
}

func (m *User) BeforeUpdate(tx *gorm.DB) error {
	if m.Password != "" {
		hashPass, err := utils.HashPassword(m.Password)
		if err != nil {
			return err
		}
		m.Password = hashPass
	}

	return nil
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	hashPass, err := utils.HashPassword(m.Password)
	if err != nil {
		return err
	}
	m.Password = hashPass
	m.Uuid = uuid.NewString()
	return nil
}
