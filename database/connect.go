package database

import (
	"github.com/Fly-Away/go-stock/models/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=B2BSprite dbname=archestock port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database. Reason : " + err.Error())
	}

	DB = database

	if migrateErr := database.AutoMigrate(&domain.User{}); migrateErr != nil {
		panic("Error migrate")
	}

}
