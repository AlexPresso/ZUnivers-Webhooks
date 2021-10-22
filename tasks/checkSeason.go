package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
	"time"
)

func checkSeason(db *gorm.DB) {
	season, err := services.FetchCurrentSeason()
	if err != nil {
		utils.Log("An error occurred while fetch current season: " + err.Error())
		return
	}

	var dbSeason structures.Season
	if res := db.Last(&dbSeason); res.Error == nil {
		season.ID = dbSeason.ID

		if (!time.Time(*dbSeason.StartDate).Equal(time.Time(*season.StartDate))) || (!time.Time(*dbSeason.EndDate).Equal(time.Time(*season.EndDate))) {
			services.DispatchEvent("new_season", dbSeason, season)
		}
	}

	db.Save(&season)
}
