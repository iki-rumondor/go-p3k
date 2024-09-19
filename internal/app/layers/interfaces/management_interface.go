package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type ManagementInterface interface {
	CreateModel(pointerModel interface{}) error

	UpdateCategory(uuid string, model *models.Category) error

	CreateShop(categoryUuid string, model *models.Shop) error
	UpdateShop(uuid string, categoryUuid string, model *models.Shop) error

	CreateProduct(userUuid string, model *models.Product) error
	UpdateProduct(userUuid string, uuid string, model *models.Product) (string, error)

	// DeleteMajor(uuid string) error
}
