package repositories

import (
	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthInterface(db *gorm.DB) interfaces.AuthInterface {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) FirstUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
