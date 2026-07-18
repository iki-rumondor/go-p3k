package models

import (
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductTransaction struct {
	ID              uint   `gorm:"primaryKey"`
	Uuid            string `gorm:"not_null;unique;size:64"`
	Quantity        int64  `gorm:"not_null"`
	IsResponse      bool   `gorm:"not_null"`
	IsAccept        bool   `gorm:"not_null"`
	IsConfirm       bool   `gorm:"not_null;default:false"`
	ProofFile       string `gorm:"not_null"`
	DeliveryProof   string `gorm:"size:255"`
	Revenue         int64  `gorm:"not_null"`
	PaymentVerified bool   `gorm:"default:false"`
	IsDelivered     bool   `gorm:"default:false"`
	IsDisbursed     bool   `gorm:"default:false"`
	DeliveredAt     int64  `gorm:"default:0"`
	CreatedAt       int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt       int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	UserID          uint   `gorm:"not_null"`
	ProductID       uint   `gorm:"not_null"`
	Product         *Product
	User            *User
}

func (m *ProductTransaction) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}

func (m *ProductTransaction) BeforeDelete(tx *gorm.DB) error {
	folder := "internal/files/transaction_proofs"
	pathFile := filepath.Join(folder, m.ProofFile)
	if err := os.Remove(pathFile); err != nil {
		log.Println(err.Error())
	}

	if m.DeliveryProof != "" {
		delFolder := "internal/files/delivery_proofs"
		delPathFile := filepath.Join(delFolder, m.DeliveryProof)
		if err := os.Remove(delPathFile); err != nil {
			log.Println(err.Error())
		}
	}
	return nil
}
