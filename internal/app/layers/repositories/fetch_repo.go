package repositories

import (
	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

type FetchRepo struct {
	db *gorm.DB
}

func NewFetchInterface(db *gorm.DB) interfaces.FetchInterface {
	return &FetchRepo{
		db: db,
	}
}

func (r *FetchRepo) GetGuests() (*[]models.Guest, error) {
	var data []models.Guest
	if err := r.db.Preload("User").Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetGuestByUuid(uuid string) (*models.Guest, error) {
	var data models.Guest
	if err := r.db.Preload("User").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCategories() (*[]models.Category, error) {
	var data []models.Category
	if err := r.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCategoryByUuid(uuid string) (*models.Category, error) {
	var data models.Category
	if err := r.db.First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetShops() (*[]models.Shop, error) {
	var data []models.Shop
	if err := r.db.Preload("User").Preload("Category").Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetShopByUuid(uuid string) (*models.Shop, error) {
	var data models.Shop
	if err := r.db.Preload("User").Preload("Category").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetAllProducts() (*[]models.Product, error) {
	var data []models.Product
	if err := r.db.Preload("Shop.Category").Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetPublicProductByUuid(uuid string) (*models.Product, error) {

	var data models.Product
	if err := r.db.Preload("Shop.Category").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetProducts(userUuid string) (*[]models.Product, error) {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var data []models.Product
	if err := r.db.Preload("Shop").Find(&data, "shop_id = ?", user.Shop.ID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetProductByUuid(userUuid, uuid string) (*models.Product, error) {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var data models.Product
	if err := r.db.Preload("Shop").First(&data, "uuid = ? AND shop_id = ?", uuid, user.Shop.ID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetProductTransactions(userUuid string) (*[]models.ProductTransaction, error) {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var data []models.ProductTransaction
	if err := r.db.Preload("Product").Find(&data, "user_id = ?", user.ID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetProductTransactionByUuid(userUuid, uuid string) (*models.ProductTransaction, error) {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var data models.ProductTransaction
	if err := r.db.Preload("Product").First(&data, "uuid = ? AND user_id = ?", uuid, user.ID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
