package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedSystemSettings(tx *gorm.DB) error {
	settings := []models.SystemSetting{
		{
			Key:   "app_name",
			Value: "P3K Aulia",
		},
		{
			Key:   "maintenance_mode",
			Value: "false",
		},
	}

	for _, setting := range settings {
		if err := tx.Create(&setting).Error; err != nil {
			return err
		}
	}
	return nil
}
