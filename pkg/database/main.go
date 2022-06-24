package database

import (
	"log"

	"github.com/Damillora/Kokomi/pkg/config"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func Initialize() {
	dsn := config.CurrentConfig.PostgresDatabase
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})


	if err != nil {
		log.Fatal("Unable to connect to database" + err.Error())
	}

	DB = db
}
