package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"github.com/iki-rumondor/go-p3k/internal/config"
	"gorm.io/gorm"
)

func SeedRoles(tx *gorm.DB) error {
	for _, role := range config.SYSTEM_ROLES {
		if err := tx.Create(&models.Role{
			Name: role,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}
