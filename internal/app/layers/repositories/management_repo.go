package repositories

import (
	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

type ManagementRepo struct {
	db *gorm.DB
}

func NewManagementInterface(db *gorm.DB) interfaces.ManagementInterface {
	return &ManagementRepo{
		db: db,
	}
}

func (r *ManagementRepo) CreateModel(modelPointer interface{}) error {
	return r.db.Create(modelPointer).Error
}

func (r *ManagementRepo) UpdateCategory(uuid string, model *models.Category) error {
	var dataDB models.Category
	if err := r.db.First(&dataDB, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	model.ID = dataDB.ID
	return r.db.Updates(model).Error
}

func (r *ManagementRepo) CreateShop(categoryUuid string, model *models.Shop) error {
	var category models.Category
	if err := r.db.First(&category, "uuid = ?", categoryUuid).Error; err != nil {
		return err
	}

	model.CategoryID = category.ID

	return r.db.Create(model).Error
}

func (r *ManagementRepo) UpdateShop(uuid string, categoryUuid string, model *models.Shop) error {
	var category models.Category
	if err := r.db.First(&category, "uuid = ?", categoryUuid).Error; err != nil {
		return err
	}

	var dataDB models.Shop
	if err := r.db.First(&dataDB, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	model.ID = dataDB.ID
	model.CategoryID = category.ID
	return r.db.Updates(model).Error
}

func (r *ManagementRepo) CreateProduct(userUuid string, model *models.Product) error {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return err
	}

	model.ShopID = user.Shop.ID
	return r.db.Create(model).Error
}

func (r *ManagementRepo) UpdateProduct(userUuid string, uuid string, model *models.Product) error {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return err
	}

	var dataDB models.Product
	if err := r.db.First(&dataDB, "uuid = ? AND shop_id = ?", uuid, user.Shop.ID).Error; err != nil {
		return err
	}

	model.ID = dataDB.ID
	return r.db.Updates(model).Error
}
