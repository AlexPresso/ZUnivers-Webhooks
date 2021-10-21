package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkItems(db *gorm.DB) {
	items, err := services.FetchItems()
	if err != nil {
		utils.Log("An error occurred while fetching items: " + err.Error())
		return
	}

	packs := checkPacks(db)
	packsMap := make(map[string]*structures.Pack)
	for _, pack := range packs {
		pack := pack
		packsMap[pack.PackID] = &pack
	}

	var dbItems []structures.Item
	db.Find(&dbItems)
	dbItemsMap := make(map[string]*structures.Item)
	for _, item := range dbItems {
		item := item
		dbItemsMap[item.ItemID] = &item
	}

	for i := 0; i < len(items); i++ {
		item := &items[i]
		dbItem := dbItemsMap[item.ItemID]

		if dbItem != nil {
			item.ID = dbItem.ID

			if utils.AreDifferent(*item, *dbItem) {
				services.DispatchEvent("item_changed", *dbItem, *item)
			}
		} else if len(dbItems) > 0 {
			services.DispatchEvent("new_item", nil, *item)
		}
	}

	db.Save(&items)
}
