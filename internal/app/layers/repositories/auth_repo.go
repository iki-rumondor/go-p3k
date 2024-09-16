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

func (r *AuthRepo) CreateModel(modelPointer interface{}) error {
	return r.db.Create(modelPointer).Error
}

func (r *AuthRepo) FirstUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepo) GetUserByUuid(uuid string) (*models.User, error) {
	var data models.User
	if err := r.db.Preload("Role").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *AuthRepo) UpdateUser(uuid string, model *models.User) error {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", uuid).Error; err != nil {
		return err
	}
	model.ID = user.ID
	return r.db.Updates(model).Error
}

func (r *AuthRepo) UnactivateUser(uuid string) error {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", uuid).Error; err != nil {
		return err
	}
	return r.db.Model(&user).Update("active", false).Error
}
