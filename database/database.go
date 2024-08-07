package database

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func Init() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("zudata.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Cannot initialize DB: " + err.Error())
	}

	err = db.AutoMigrate(
		&structures.Config{},
		&structures.Patchnote{},
		&structures.Status{},
		&structures.Pack{},
		&structures.Item{},
		&structures.Season{},
		&structures.Banner{},
		&structures.Event{},
		&structures.AchievementCategory{},
		&structures.Achievement{},
		&structures.Challenge{},
		&structures.ShopEntry{},
		&structures.JsonResponseSpec{},
	)

	if err != nil {
		log.Fatal("Error migrating schema: " + err.Error())
	}
	return
}
