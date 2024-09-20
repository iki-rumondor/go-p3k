package repositories

import (
	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
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

func (r *ManagementRepo) CheckUniqueNik(nik string) bool {
	rows := r.db.First(&models.Citizen{}, "nik = ?", nik).RowsAffected
	return rows == 0
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

func (r *ManagementRepo) UpdateCitizen(uuid string, model *models.Citizen) error {
	var dataDB models.Citizen
	if err := r.db.First(&dataDB, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	if dataDB.Nik != model.Nik {
		if rows := r.db.First(&models.Citizen{}, "nik = ? AND id != ?", model.Nik, dataDB.ID).RowsAffected; rows == 1 {
			return response.BADREQ_ERR("Nik yang digunakan sudah terdaftar")
		}

		model.User = &models.User{
			Username: model.Nik,
			Password: model.Nik,
		}
	}

	model.ID = dataDB.ID
	return r.db.Updates(model).Error
}

func (r *ManagementRepo) UpdateMember(uuid string, model *models.Member) error {
	var dataDB models.Member
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

func (r *ManagementRepo) UpdateProduct(userUuid string, uuid string, model *models.Product) (string, error) {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return "", err
	}

	var dataDB models.Product
	if err := r.db.First(&dataDB, "uuid = ? AND shop_id = ?", uuid, user.Shop.ID).Error; err != nil {
		return "", err
	}

	model.ID = dataDB.ID
	if err := r.db.Updates(model).Error; err != nil {
		return "", err
	}

	return dataDB.Image, nil
}

func (r *ManagementRepo) CreateActivity(userUuid string, model *models.Activity) error {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return err
	}

	model.CreatedUserID = user.ID
	model.UpdatedUserID = user.ID
	return r.db.Create(model).Error
}

func (r *ManagementRepo) UpdateActivity(userUuid string, uuid string, model *models.Activity) (string, error) {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return "", err
	}

	var dataDB models.Activity
	if err := r.db.First(&dataDB, "uuid = ?", uuid).Error; err != nil {
		return "", err
	}

	model.ID = dataDB.ID
	model.UpdatedUserID = user.ID
	if err := r.db.Updates(model).Error; err != nil {
		return "", err
	}

	return dataDB.ImageName, nil
}

func (r *ManagementRepo) DeleteActivity(uuid string) (string, error) {
	var dataDB models.Activity
	if err := r.db.First(&dataDB, "uuid = ?", uuid).Error; err != nil {
		return "", err
	}

	if err := r.db.Delete(&dataDB).Error; err != nil {
		return "", err
	}

	return dataDB.ImageName, nil
}

func (r *ManagementRepo) GetUserByUuid(userUuid string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *ManagementRepo) CheckExistMemberActivity(memberUuid, activityUuid string) (bool, error) {
	var member models.Member
	if err := r.db.First(&member, "uuid = ?", memberUuid).Error; err != nil {
		return false, err
	}

	var activity models.Activity
	if err := r.db.First(&activity, "uuid = ?", activityUuid).Error; err != nil {
		return false, err
	}

	rows := r.db.First(&models.MemberActivity{}, "member_id = ? AND activity_id = ?", member.ID, activity.ID).RowsAffected

	return rows == 1, nil
}

func (r *ManagementRepo) CreateMemberActivity(userID uint, memberUuid, activityUuid string) error {
	var member models.Member
	if err := r.db.First(&member, "uuid = ?", memberUuid).Error; err != nil {
		return err
	}

	var activity models.Activity
	if err := r.db.First(&activity, "uuid = ?", activityUuid).Error; err != nil {
		return err
	}

	model := models.MemberActivity{
		MemberID:      member.ID,
		ActivityID:    activity.ID,
		CreatedUserID: userID,
	}

	return r.db.Create(&model).Error
}

func (r *ManagementRepo) DeleteMemberActivity(memberUuid, activityUuid string) error {
	var member models.Member
	if err := r.db.First(&member, "uuid = ?", memberUuid).Error; err != nil {
		return err
	}

	var activity models.Activity
	if err := r.db.First(&activity, "uuid = ?", activityUuid).Error; err != nil {
		return err
	}

	var model models.MemberActivity
	if err := r.db.First(&model, "member_id = ? AND activity_id = ?", member.ID, activity.ID).Error; err != nil {
		return err
	}

	return r.db.Delete(&model).Error
}
