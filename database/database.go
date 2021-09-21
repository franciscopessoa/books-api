package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	db := database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(18)
	config.SetMaxIdleConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetInstance() *gorm.DB {
	return db
}
