package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const NewSeasonEvent = "new_season"

func checkSeason(db *gorm.DB, embeds *[]discord.Embed) {
	if utils.EventsAllDisabled([]string{NewSeasonEvent}) {
		return
	}

	season, resSpec, err := services.FetchCurrentSeason()
	if err != nil {
		utils.Log("An error occurred while fetching current season: " + err.Error())
		return
	}

	checkResponse(db, embeds, resSpec)

	var dbSeason structures.Season
	if res := db.Last(&dbSeason); res.Error == nil {
		season.ID = dbSeason.ID

		if utils.AreDifferent(dbSeason, season) {
			services.MakeEmbed(NewSeasonEvent, dbSeason, season, embeds)
		}
	}

	db.Save(&season)
}
