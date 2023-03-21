package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresConnection(config Config) (error, *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s", config.Host, config.User, config.Password, config.Name, config.Port, config.TimeZone)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database " + config.Host)
		return err, nil
	}

	setup(connection)

	return err, connection
}

func setup(db *gorm.DB) {
	db.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
}
