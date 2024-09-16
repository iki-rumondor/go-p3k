package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type AuthInterface interface {
	GetUserByUuid(uuid string) (*models.User, error)
	FirstUserByUsername(username string) (*models.User, error)
	CreateModel(modelPointer interface{}) error
	UpdateUser(uuid string, model *models.User) error
	UnactivateUser(uuid string) error
}
