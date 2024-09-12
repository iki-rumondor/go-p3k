package config

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

var SYSTEM_ROLES = []string{"ADMIN", "DEPARTMENT"}
var ADMIN_USER = models.User{
	Name:     "Administrator",
	Username: "admin",
	Password: "admin123",
	Active:   true,
	RoleID:   1,
}
