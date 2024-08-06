package tasks

import (
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const ShopChangedEvent = "shop_changed"

func checkShopEntries(db *gorm.DB, embeds *[]discord.Embed) {
	if !utils.EventsEnabled([]string{ShopChangedEvent}) {
		return
	}

	entries, resSpec, err := services.FetchShop()
	if err != nil {
		utils.Log("An error occurred while fetching shop entries: " + err.Error())
		return
	}

	checkResponse(db, embeds, ShopChangedEvent, resSpec)

	var dbEntries []structures.ShopEntry
	db.Find(&dbEntries)
	dbEntriesMap := make(map[string]*structures.ShopEntry)
	for _, entry := range dbEntries {
		entry := entry
		dbEntriesMap[entry.ShopEntryID] = &entry
	}

	difference := false
	for i := 0; i < len(entries); i++ {
		entry := &entries[i]

		if len(dbEntries) > i {
			entry.ID = (&dbEntries[i]).ID
		}

		difference = dbEntriesMap[entry.ShopEntryID] == nil
	}

	if difference {
		*embeds = append(*embeds, *makeShopEmbed(entries))
	}

	db.Save(entries)
}

func makeShopEmbed(entries []structures.ShopEntry) *discord.Embed {
	frontBaseUrl := viper.GetString("frontBaseUrl")

	embed := services.DefaultEmbed(ShopChangedEvent, fmt.Sprintf("%s/echoppe", frontBaseUrl))
	itemsField := &discord.EmbedField{
		Name:   "Items",
		Value:  "",
		Inline: true,
	}

	itemPrices := make(map[string]int)
	maxNameLength := 0

	for _, entry := range entries {
		shiny := ""

		switch entry.ShopInventory.ShinyLevel {
		case 0:
			shiny = ""
			break
		case 1:
			shiny = " (dorée)"
			break
		case 2:
			shiny = " (chroma)"
			break
		}

		name := fmt.Sprintf("%s%s", entry.ShopInventory.Item.Name, shiny)
		itemPrices[name] = entry.ShopInventory.Balance

		if len(name) > maxNameLength {
			maxNameLength = len(name)
		}
	}

	emoji := viper.GetString("emojis.balance")
	for name, price := range itemPrices {
		itemsField.Value += fmt.Sprintf("`%-*s %d`%s\n", maxNameLength, name, price, emoji)
	}

	embed.Fields = []*discord.EmbedField{itemsField}

	return embed
}
