package database

import (
	"gorm.io/gorm"
	"jan-galek/api/src/entities"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(entities.Article{})
	seed(db)
}

func seed(db *gorm.DB) {
	articles := []entities.Article{
		{Title: "General", Link: "/general"},
		{Title: "Off-Topic", Link: "/off-topic"},
		{Title: "Suggestions", Link: "/suggestions"},
	}

	for _, article := range articles {
		db.Create(&article)
	}
}
