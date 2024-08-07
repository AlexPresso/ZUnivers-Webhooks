package tasks

import (
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const ResponseChangedEvent = "response_changed"

func checkResponse(db *gorm.DB, embeds *[]discord.Embed, resSpec structures.JsonResponseSpec) {
	if utils.EventsAllDisabled([]string{ResponseChangedEvent}) {
		return
	}

	var prevSpec structures.JsonResponseSpec
	if res := db.Where("endpoint_uri = ?", resSpec.EndpointURI).Find(&prevSpec); res.Error == nil {
		db.Save(&resSpec)
		if prevSpec.EndpointURI == "" {
			return
		}

		diff, hasDiff := utils.GenerateDiff(prevSpec.Value, resSpec.Value)
		if !hasDiff {
			return
		}

		embed := services.DefaultEmbed(ResponseChangedEvent, resSpec.EndpointURI)
		embed.Description += fmt.Sprintf("\n```diff\n%s```", diff)

		*embeds = append(*embeds, *embed)
	}
}
