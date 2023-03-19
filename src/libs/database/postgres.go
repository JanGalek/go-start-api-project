package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jan-galek/api/src/libs"
)

func GetConnection() (error, *gorm.DB) {
	dsn := "host=" + libs.GetEnv("DATABASE_HOST") +
		" user=" + libs.GetEnv("DATABASE_USER") +
		" password=" + libs.GetEnv("DATABASE_PASSWORD") +
		" dbname=" + libs.GetEnv("DATABASE_NAME") +
		" port=" + libs.GetEnv("DATABASE_PORT") +
		" sslmode=" + libs.GetEnv("DATABASE_SSLMODE") +
		" TimeZone=" + libs.GetEnv("DATABASE_TIME_ZONE")
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database " + libs.GetEnv("DATABASE_HOST"))
		return err, nil
	}

	return err, connection
}
