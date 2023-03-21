package database

import (
	"fmt"
	"jan-galek/api/src/libs"
)

type Config struct {
	Host     string
	Name     string
	Port     string
	User     string
	Password string
	TimeZone string
}

func (config *Config) GetPostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s", config.Host, config.User, config.Password, config.Name, config.Port, config.TimeZone)
}
func (config *Config) GetPostgresDSNMigration() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable" /*config.Host*/, "localhost", config.User, config.Password, config.Name, config.Port)
}

func (config *Config) CreatePostgres() Config {
	return Config{
		Host:     libs.GetEnv("DATABASE_HOST"),
		User:     libs.GetEnv("DATABASE_USER"),
		Password: libs.GetEnv("DATABASE_PASSWORD"),
		Name:     libs.GetEnv("DATABASE_NAME"),
		Port:     libs.GetEnv("DATABASE_PORT"),
		TimeZone: libs.GetEnv("DATABASE_TIME_ZONE"),
	}
}
