package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const NewPackEvent = "new_pack"
const PackChangedEvent = "pack_changed"
const PackRemovedEvent = "pack_removed"

func checkPacks(db *gorm.DB, embeds *[]discord.Embed) (packs []structures.Pack) {
	if !utils.EventsEnabled([]string{NewPackEvent, PackChangedEvent, PackRemovedEvent}) {
		return
	}

	packs, resSpec, err := services.FetchPacks()
	if err != nil {
		utils.Log("An error occurred while fetching packs: " + err.Error())
		return
	}

	checkResponse(db, embeds, resSpec)

	packsMap := make(map[string]*structures.Pack)
	for _, pack := range packs {
		pack := pack
		packsMap[pack.PackID] = &pack
	}

	var dbPacks []structures.Pack
	db.Find(&dbPacks)
	dbPacksMap := make(map[string]*structures.Pack)
	for _, pack := range dbPacks {
		pack := pack
		dbPacksMap[pack.PackID] = &pack
	}

	for i := 0; i < len(packs); i++ {
		pack := &packs[i]
		dbPack := dbPacksMap[pack.PackID]

		if dbPack != nil {
			pack.ID = dbPack.ID

			if utils.AreDifferent(*dbPack, *pack) {
				*embeds = append(*embeds, *services.MakeEmbed(PackChangedEvent, *dbPack, *pack))
			}
		} else if len(dbPacksMap) > 0 {
			*embeds = append(*embeds, *services.MakeEmbed(NewPackEvent, nil, *pack))
		}
	}

	db.Save(&packs)

	for _, pack := range dbPacks {
		pack := pack
		if packsMap[pack.PackID] == nil {
			db.Delete(&pack)
			*embeds = append(*embeds, *services.MakeEmbed(PackRemovedEvent, nil, pack))
		}
	}

	return
}
