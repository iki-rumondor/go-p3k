package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedCategories(tx *gorm.DB) error {
	categories := []models.Category{
		{Name: "Sembako"},
		{Name: "Makanan"},
		{Name: "Minuman"},
		{Name: "Pakaian"},
		{Name: "Jasa"},
	}

	for _, category := range categories {
		if err := tx.Create(&category).Error; err != nil {
			return err
		}
	}
	return nil
}
