package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedRegions(tx *gorm.DB) error {
	regions := []models.Region{
		{Name: "SULAWESI UTARA"},
		{Name: "SULAWESI TENGAH"},
		{Name: "SULAWESI SELATAN"},
	}

	for _, region := range regions {
		if err := tx.Create(&region).Error; err != nil {
			return err
		}
	}
	return nil
}
