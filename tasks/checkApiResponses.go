package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkApiResponses(db *gorm.DB, embeds *[]discord.Embed) {
	if utils.EventsAllDisabled([]string{ResponseChangedEvent}) {
		return
	}

	routes := [...]string{
		"/public/lucky/jackpot",
		"/public/tournament/latest",
		"/public/tower/stats",
		"/public/tower/alexpresso",
		"/public/corporation",
		"/public/market",
		"/public/user/alexpresso/overview",
		"/public/user/alexpresso/activity",
		"/public/evolution/alexpresso",
		"/public/loot/alexpresso?year=0",
	}

	for _, route := range routes {
		_, resSpec, err := services.FetchRoute(route)
		if err != nil {
			utils.Log("An error occurred while fetching " + route + " route: " + err.Error())
			continue
		}

		checkResponse(db, embeds, resSpec)
	}
}
