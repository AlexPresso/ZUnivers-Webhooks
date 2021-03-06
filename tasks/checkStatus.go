package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkStatus(db *gorm.DB, embeds *[]discord.Embed) {
	status, err := services.FetchStatus()
	if err != nil {
		utils.Log("An error occurred while fetching status: " + err.Error())
		return
	}

	var currStatus structures.Status
	if res := db.First(&currStatus); res.Error == nil {
		status.ID = currStatus.ID

		if utils.AreDifferent(currStatus, status) {
			*embeds = append(*embeds, *services.MakeEmbed("status_changed", currStatus, status))
		}
	}

	db.Save(&status)
}
