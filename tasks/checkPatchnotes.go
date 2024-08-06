package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
	"time"
)

const NewPatchnoteEvent = "new_patchnote"

func checkPatchnotes(db *gorm.DB, embeds *[]discord.Embed) {
	if !utils.EventsEnabled([]string{NewPatchnoteEvent}) {
		return
	}

	patchnotes, resSpec, err := services.FetchPatchnotes()
	if err != nil {
		utils.Log("An error occurred while fetching patch notes: " + err.Error())
		return
	}

	checkResponse(db, embeds, NewPatchnoteEvent, resSpec)

	if len(patchnotes) == 0 {
		return
	}

	var latestPatchnote structures.Patchnote
	if res := db.Last(&latestPatchnote); res.Error == nil {
		if time.Time(*latestPatchnote.Date).Before(time.Time(*patchnotes[0].Date)) {
			*embeds = append(*embeds, *services.MakeEmbed(NewPatchnoteEvent, nil, patchnotes[0]))
			db.Save(&patchnotes[0])
		}
	} else {
		db.Save(&patchnotes[0])
	}
}
