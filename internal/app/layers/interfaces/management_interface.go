package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type ManagementInterface interface {
	CheckUniqueNik(nik string) bool
	CreateModel(pointerModel interface{}) error

	UpdateCategory(uuid string, model *models.Category) error
	UpdateCitizen(uuid string, model *models.Citizen) error
	UpdateMember(uuid string, model *models.Member) error

	CreateShop(categoryUuid string, model *models.Shop) error
	UpdateShop(uuid string, categoryUuid string, model *models.Shop) error

	CreateProduct(userUuid string, model *models.Product) error
	UpdateProduct(userUuid string, uuid string, model *models.Product) (string, error)
}
