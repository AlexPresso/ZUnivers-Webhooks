package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkPacks(db *gorm.DB) (packs []structures.Pack) {
	packs, err := services.FetchPacks()
	if err != nil {
		utils.Log("An error occurred while fetching packs: " + err.Error())
		return
	}

	var storedPacks []structures.Pack
	db.Find(&storedPacks)
	packsMap := make(map[string]*structures.Pack)
	for _, pack := range storedPacks {
		pack := pack
		packsMap[pack.PackID] = &pack
	}

	for i := 0; i < len(packs); i++ {
		pack := &packs[i]
		dbPack := packsMap[pack.PackID]

		if dbPack != nil {
			pack.ID = dbPack.ID

			if utils.AreDifferent(*dbPack, *pack) {
				services.DispatchEvent("pack_changed", *dbPack, *pack)
			}
		} else if len(packsMap) > 0 {
			services.DispatchEvent("new_pack", nil, *pack)
		}
	}

	db.Save(&packs)
	return
}
