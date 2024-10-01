package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type ManagementInterface interface {
	CheckUniqueNik(nik string) bool
	CreateModel(pointerModel interface{}) error

	UpdateCategory(uuid string, model *models.Category) error
	UpdateCitizen(uuid string, model *models.Citizen) error
	UpdateMember(uuid string, model *models.Member) error

	// CreateShop(categoryUuid string, model *models.Shop) error
	// UpdateShop(uuid string, categoryUuid string, model *models.Shop) error

	CreateProduct(userUuid string, model *models.Product) error
	UpdateProduct(userUuid string, uuid string, model *models.Product) (string, error)

	CreateActivity(userUuid string, model *models.Activity) error
	UpdateActivity(userUuid string, uuid string, model *models.Activity) (string, error)
	DeleteActivity(uuid string) (string, error)

	GetUserByUuid(uuid string) (*models.User, error)
	CheckExistMemberActivity(memberUuid, activityUuid string) (bool, error)

	CreateMemberActivity(userID uint, memberUuid, activityUuid string) error
	DeleteMemberActivity(memberUuid, activityUuid string) error
}
