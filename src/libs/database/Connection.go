package database

import (
	"gorm.io/gorm"
)

func GetConnection() (error, *gorm.DB) {
	config := Config{}
	return GetPostgresConnection(config.CreatePostgres())
}
