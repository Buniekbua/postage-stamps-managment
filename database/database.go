package database

import (
	"log"
	"os"

	"github.com/Buniekbua/postage-stamps-managment/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbRepository struct {
	Db *gorm.DB
}

var Database DbRepository

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("dbstamps.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!\n", err.Error())
		os.Exit(1)
	}
	log.Println("Successfully connected to the database!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(&models.User{}, &models.Stamp{})

	Database = DbRepository{Db: db}
}
