package database

import (
	"log"
	"taskproject/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dburl      = "./database.sqlite"
	dbInstance *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance, nil
	}

	db, err := gorm.Open(sqlite.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	dbInstance = db
	return dbInstance, nil
}
