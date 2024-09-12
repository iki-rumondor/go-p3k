package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type AuthInterface interface {
	FirstUserByUsername(username string) (*models.User, error)
}
