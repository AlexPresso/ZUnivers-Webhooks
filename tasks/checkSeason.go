package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkSeason(db *gorm.DB, embeds *[]discord.Embed) {
	season, err := services.FetchCurrentSeason()
	if err != nil {
		utils.Log("An error occurred while fetching current season: " + err.Error())
		return
	}

	var dbSeason structures.Season
	if res := db.Last(&dbSeason); res.Error == nil {
		season.ID = dbSeason.ID

		if utils.AreDifferent(dbSeason, season) {
			*embeds = append(*embeds, *services.MakeEmbed("new_season", dbSeason, season))
		}
	}

	db.Save(&season)
}
