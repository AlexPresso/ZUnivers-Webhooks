package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const ResponseChangedEvent = "response_changed"

func checkResponse(db *gorm.DB, embeds *[]discord.Embed, event string, resSpec map[string]interface{}) {
	if !utils.EventsEnabled([]string{ResponseChangedEvent}) {
		return
	}

}
