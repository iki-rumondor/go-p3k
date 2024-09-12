package migrate

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type Model struct {
	Model interface{}
}

func GetAllModels() []Model {
	return []Model{
		{Model: models.Role{}},
		{Model: models.User{}},
	}
}
