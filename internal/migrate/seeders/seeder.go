package seeders

import (
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := SeedRoles(tx); err != nil {
			return err
		}
		if err := SeedRegions(tx); err != nil {
			return err
		}
		if err := SeedUsers(tx); err != nil {
			return err
		}
		if err := SeedCategories(tx); err != nil {
			return err
		}
		if err := SeedProducts(tx); err != nil {
			return err
		}
		if err := SeedTransactions(tx); err != nil {
			return err
		}
		if err := SeedActivities(tx); err != nil {
			return err
		}
		if err := SeedMemberActivities(tx); err != nil {
			return err
		}
		if err := SeedSystemSettings(tx); err != nil {
			return err
		}
		if err := SeedTutorials(tx); err != nil {
			return err
		}
		return nil
	})
}
