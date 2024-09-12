package config

import (
	"github.com/iki-rumondor/go-p3k/internal/app/layers/handlers"
	"github.com/iki-rumondor/go-p3k/internal/app/layers/repositories"
	"github.com/iki-rumondor/go-p3k/internal/app/layers/services"
	"gorm.io/gorm"
)

type Handlers struct {
	AuthHandler *handlers.AuthHandler
}

func GetAppHandlers(db *gorm.DB) *Handlers {

	auth_repo := repositories.NewAuthInterface(db)
	auth_service := services.NewAuthService(auth_repo)
	auth_handler := handlers.NewAuthHandler(auth_service)

	return &Handlers{
		AuthHandler: auth_handler,
	}
}
