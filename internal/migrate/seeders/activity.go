package seeders

import (
	"time"

	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedActivities(tx *gorm.DB) error {
	var adminUser models.User
	if err := tx.First(&adminUser, "role_id = ?", 1).Error; err != nil {
		return err
	}

	activities := []models.Activity{
		{
			Group:         1,
			Title:         "Kerja Bakti Bersama",
			Description:   "Membersihkan saluran air lingkungan dan halaman balai desa.",
			Location:      "Balai Desa Sentosa",
			ImageName:     "default_activity.png",
			CreatedUserID: adminUser.ID,
			UpdatedUserID: adminUser.ID,
			StartTime:     time.Now().UnixMilli(),
			EndTime:       time.Now().Add(2 * time.Hour).UnixMilli(),
		},
	}

	for _, activity := range activities {
		if err := tx.Create(&activity).Error; err != nil {
			return err
		}
	}
	return nil
}
