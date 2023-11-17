package initializers

import "GO_AUTH/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}