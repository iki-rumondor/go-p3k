package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedTransactions(tx *gorm.DB) error {
	var user models.User
	if err := tx.First(&user, "role_id = ?", 5).Error; err != nil {
		return err
	}

	var product models.Product
	if err := tx.First(&product).Error; err != nil {
		return err
	}

	transactions := []models.ProductTransaction{
		{
			Quantity:        2,
			IsResponse:      true,
			IsAccept:        true,
			IsConfirm:       true,
			ProofFile:       "default_proof.png",
			Revenue:         product.Price * 2,
			PaymentVerified: true,
			IsDelivered:     true,
			UserID:          user.ID,
			ProductID:       product.ID,
		},
	}

	for _, transaction := range transactions {
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
	}
	return nil
}
