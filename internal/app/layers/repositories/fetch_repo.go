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
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var productIDs []uint
	if err := r.db.Model(&models.Product{}).Where("shop_id = ?", user.Shop.ID).Pluck("id", &productIDs).Error; err != nil {
		return nil, err
	}

	var data models.ProductTransaction
	if err := r.db.Preload("Product").Preload("User.Role").Preload("User.Guest").First(&data, "uuid = ? AND product_id IN (?)", uuid, productIDs).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetProductTransactionsByShop(userUuid string) (*[]models.ProductTransaction, error) {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var productIDs []uint
	if err := r.db.Model(&models.Product{}).Where("shop_id = ?", user.Shop.ID).Pluck("id", &productIDs).Error; err != nil {
		return nil, err
	}

	var data []models.ProductTransaction
	if err := r.db.Preload("Product").Preload("User").Find(&data, "product_id IN (?)", productIDs).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCitizens() (*[]models.Citizen, error) {
	var data []models.Citizen
	if err := r.db.Preload("User").Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCitizenByUuid(uuid string) (*models.Citizen, error) {
	var data models.Citizen
	if err := r.db.Preload("User").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetMembers() (*[]models.Member, error) {
	var data []models.Member
	if err := r.db.Preload("User").Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetMemberByUuid(uuid string) (*models.Member, error) {
	var data models.Member
	if err := r.db.Preload("User").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetActivities() (*[]models.Activity, error) {
	var data []models.Activity
	if err := r.db.Preload("CreatedUser").Preload("UpdatedUser").Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetActivityByUuid(uuid string) (*models.Activity, error) {
	var data models.Activity
	if err := r.db.Preload("CreatedUser").Preload("UpdatedUser").Preload("Members.Member").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetMembersNotInActivity(activityUuid string) (*[]models.Member, error) {
	var activity models.Activity
	if err := r.db.First(&activity, "uuid = ?", activityUuid).Error; err != nil {
		return nil, err
	}

	var memberIDs []uint
	if err := r.db.Model(&models.MemberActivity{}).Where("activity_id = ?", activity.ID).Pluck("member_id", &memberIDs).Error; err != nil {
		return nil, err
	}

	var resp []models.Member
	if len(memberIDs) == 0 {
		if err := r.db.Find(&resp).Error; err != nil {
			return nil, err
		}
		return &resp, nil
	}

	if err := r.db.Find(&resp, "id NOT IN (?)", memberIDs).Error; err != nil {
		return nil, err
	}
	return &resp, nil
}
