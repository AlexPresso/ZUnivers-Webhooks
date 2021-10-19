package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"zunivers-webhooks/structures"
)

func Init() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("zudata.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Cannot initialize DB: " + err.Error())
	}

	if err := db.AutoMigrate(&structures.Config{}, &structures.Patchnote{}, &structures.Status{}); err != nil {
		log.Fatal("Error migrating schema: " + err.Error())
	}
	return
}
