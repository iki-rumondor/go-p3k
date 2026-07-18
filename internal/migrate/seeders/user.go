package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedUsers(tx *gorm.DB) error {
	// 1. Seed ADMIN User
	adminUser := models.User{
		Name:     "Administrator",
		Username: "admin",
		Password: "admin123",
		Active:   true,
		RoleID:   1,
	}
	if err := tx.Create(&adminUser).Error; err != nil {
		return err
	}

	// 2. Seed MEMBER User & Profile
	memberProfile := models.Member{
		Name:        "Member Default",
		Position:    "Staf",
		IsImportant: true,
		IsHeadgroup: false,
		Group:       1,
		User: &models.User{
			Name:     "Member Default",
			Username: "member",
			Password: "member123",
			Active:   true,
			RoleID:   2,
		},
	}
	if err := tx.Create(&memberProfile).Error; err != nil {
		return err
	}

	// 3. Seed UMKM User & Shop Profile
	shopProfile := models.Shop{
		Name:          "Toko UMKM",
		Owner:         "UMKM Owner",
		Address:       "Jl. UMKM No. 1",
		PhoneNumber:   "081234567890",
		ShopImage:     "default_shop.png",
		IdentityImage: "default_identity.png",
		User: &models.User{
			Name:     "UMKM Owner",
			Username: "umkm",
			Password: "umkm123",
			Active:   true,
			RoleID:   3,
		},
	}
	if err := tx.Create(&shopProfile).Error; err != nil {
		return err
	}

	// 4. Seed GUEST User & Guest Profile
	guestProfile := models.Guest{
		Name:        "Guest Default",
		Address:     "Jl. Guest No. 1",
		PhoneNumber: "081234567891",
		User: &models.User{
			Name:     "Guest Default",
			Username: "guest",
			Password: "guest123",
			Active:   true,
			RoleID:   4,
		},
	}
	if err := tx.Create(&guestProfile).Error; err != nil {
		return err
	}

	// 5. Seed CITIZEN User & Citizen Profile
	var firstRegion models.Region
	if err := tx.First(&firstRegion).Error; err != nil {
		return err
	}

	citizenProfile := models.Citizen{
		Name:        "Citizen Default",
		Nik:         "1234567890123456",
		Address:     "Jl. Citizen No. 1",
		PhoneNumber: "081234567892",
		RegionID:    firstRegion.ID,
		User: &models.User{
			Name:     "Citizen Default",
			Username: "citizen",
			Password: "citizen123",
			Active:   true,
			RoleID:   5,
		},
	}
	if err := tx.Create(&citizenProfile).Error; err != nil {
		return err
	}

	return nil
}
