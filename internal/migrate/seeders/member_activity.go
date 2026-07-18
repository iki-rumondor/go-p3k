package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedMemberActivities(tx *gorm.DB) error {
	var member models.Member
	if err := tx.First(&member).Error; err != nil {
		return err
	}

	var activity models.Activity
	if err := tx.First(&activity).Error; err != nil {
		return err
	}

	memberActivities := []models.MemberActivity{
		{
			MemberID:        member.ID,
			ActivityID:      activity.ID,
			AttendenceImage: "default_attendance.png",
			IsAccept:        true,
		},
	}

	for _, ma := range memberActivities {
		if err := tx.Create(&ma).Error; err != nil {
			return err
		}
	}
	return nil
}
