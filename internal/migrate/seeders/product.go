package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedProducts(tx *gorm.DB) error {
	var shop models.Shop
	if err := tx.First(&shop).Error; err != nil {
		return err
	}

	var category models.Category
	if err := tx.First(&category).Error; err != nil {
		return err
	}

	products := []models.Product{
		{
			Name:       "Beras Premium 5kg",
			Price:      65000,
			Stock:      100,
			Unit:       "karung",
			Image:      "beras.png",
			ShopID:     shop.ID,
			CategoryID: category.ID,
		},
		{
			Name:       "Minyak Goreng 2L",
			Price:      32000,
			Stock:      50,
			Unit:       "botol",
			Image:      "minyak.png",
			ShopID:     shop.ID,
			CategoryID: category.ID,
		},
	}

	for _, product := range products {
		if err := tx.Create(&product).Error; err != nil {
			return err
		}
	}
	return nil
}
