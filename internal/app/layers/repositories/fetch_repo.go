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

func (r *FetchRepo) GetGuestByUser(userUuid string) (*models.Guest, error) {
	var user models.User
	if err := r.db.Preload("Guest").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var data models.Guest
	if err := r.db.First(&data, "id = ?", user.Guest.ID).Error; err != nil {
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

func (r *FetchRepo) GetShops(limit int) (*[]models.Shop, error) {
	var data []models.Shop
	query := r.db.Preload("User")
	if limit != 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetShopByUuid(uuid string) (*models.Shop, error) {
	var data models.Shop
	if err := r.db.Preload("User").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetShopByUser(userUuid string) (*models.Shop, error) {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var data models.Shop
	if err := r.db.First(&data, "id = ?", user.Shop.ID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetAllProducts(limit int, shopUuid string) (*[]models.Product, error) {
	var data []models.Product
	query := r.db.Preload("Shop").Preload("Category").Limit(limit)
	if shopUuid != "" {
		var shop models.Shop
		if err := r.db.First(&shop, "uuid = ?", shopUuid).Error; err != nil {
			return nil, err
		}
		query = query.Where("shop_id = ?", shop.ID)
	}

	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetPublicProductByUuid(uuid string) (*models.Product, error) {

	var data models.Product
	if err := r.db.Preload("Shop").Preload("Category").First(&data, "uuid = ?", uuid).Error; err != nil {
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
	if err := r.db.Preload("Shop").Preload("Category").Find(&data, "shop_id = ?", user.Shop.ID).Error; err != nil {
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
	if err := r.db.Preload("Shop").Preload("Category").First(&data, "uuid = ? AND shop_id = ?", uuid, user.Shop.ID).Error; err != nil {
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

func (r *FetchRepo) GetProductTransactionsByShop(userUuid string, isAccept bool) (*[]models.ProductTransaction, error) {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var productIDs []uint
	if err := r.db.Model(&models.Product{}).Where("shop_id = ?", user.Shop.ID).Pluck("id", &productIDs).Error; err != nil {
		return nil, err
	}

	var data []models.ProductTransaction
	if err := r.db.Preload("Product").Preload("User").Find(&data, "product_id IN (?) AND is_accept = ?", productIDs, isAccept).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCitizens() (*[]models.Citizen, error) {
	var data []models.Citizen
	if err := r.db.Preload("User").Preload("Region").Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCitizensWithRegion(regionId string) (*[]models.Citizen, error) {
	var data []models.Citizen
	query := r.db.Preload("User").Preload("Region")
	if regionId != "" {
		query = query.Where("region_id = ?", regionId)
	}
	
	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetCitizenByUuid(uuid string) (*models.Citizen, error) {
	var data models.Citizen
	if err := r.db.Preload("User").Preload("Region").First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetMembers(group string) (*[]models.Member, error) {
	var data []models.Member
	query := r.db.Preload("User")
	if group != "" {
		query = query.Where("`group` = ?", group)
	}
	if err := query.Find(&data).Error; err != nil {
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

func (r *FetchRepo) GetMemberByUserUuid(userUuid string) (*models.Member, error) {
	var user models.User
	if err := r.db.Preload("Member").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	return user.Member, nil
}

func (r *FetchRepo) GetActivities(limit int, group string) (*[]models.Activity, error) {
	var data []models.Activity
	query := r.db.Preload("CreatedUser").Preload("UpdatedUser").Preload("Members.Member").Limit(limit)
	if group != "" {
		query = query.Where("`group` = ?", group)
	}

	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *FetchRepo) GetActivitiesWithDate(limit int, group, startDate, endDate string) (*[]models.Activity, error) {
	var data []models.Activity
	query := r.db.Preload("CreatedUser").Preload("UpdatedUser").Preload("Members.Member").Limit(limit)
	if group != "" {
		query = query.Where("`group` = ?", group)
	}

	if startDate != "" {
		query = query.Where("start_time >= ?", startDate)
	}

	if endDate != "" {
		query = query.Where("end_time <= ?", endDate)
	}

	if err := query.Find(&data).Error; err != nil {
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

func (r *FetchRepo) GetMemberActivity(userUuid, activityUuid string) (*models.MemberActivity, error) {
	var user models.User
	if err := r.db.Preload("Member").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var activity models.Activity
	if err := r.db.First(&activity, "uuid = ?", activityUuid).Error; err != nil {
		return nil, err
	}

	var result models.MemberActivity
	if err := r.db.Preload("Activity.UpdatedUser").First(&result, "member_id = ? AND activity_id = ?", user.Member.ID, activity.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *FetchRepo) GetMemberActivities() (*[]models.MemberActivity, error) {

	var result []models.MemberActivity
	if err := r.db.Preload("Activity.UpdatedUser").Preload("Member").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *FetchRepo) CountGuestsInactive() (int64, error) {
	var count int64
	if err := r.db.Joins("User").Model(&models.Guest{}).Where("User.active = ?", false).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *FetchRepo) CountShopsInactive() (int64, error) {
	var count int64
	if err := r.db.Joins("User").Model(&models.Shop{}).Where("User.active = ?", false).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *FetchRepo) CountShopProducts(userUuid string) (int64, error) {
	var count int64

	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return count, err
	}

	if err := r.db.Model(&models.Product{}).Where("shop_id = ?", user.Shop.ID).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *FetchRepo) CountShopUnprocessTransactions(userUuid string) (int64, error) {
	var count int64

	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return count, err
	}

	var productIDs []uint
	if err := r.db.Model(&models.Product{}).Where("shop_id = ?", user.Shop.ID).Pluck("id", &productIDs).Error; err != nil {
		return count, err
	}

	if err := r.db.Model(&models.ProductTransaction{}).Where("product_id IN (?) AND is_response = ?", productIDs, false).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *FetchRepo) CountUserSuccessTransactions(userUuid string) (int64, error) {
	var count int64

	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return count, err
	}

	if err := r.db.Model(&models.ProductTransaction{}).Where("user_id = ? AND is_accept = ?", user.ID, true).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *FetchRepo) CountUserUnprocessTransactions(userUuid string) (int64, error) {
	var count int64

	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return count, err
	}

	if err := r.db.Model(&models.ProductTransaction{}).Where("user_id = ? AND is_response = ?", user.ID, false).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (r *FetchRepo) CountActivities() (int64, error) {
	var count int64

	if err := r.db.Model(&models.Activity{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
