package tasks

import (
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
)

func checkShopEntries(db *gorm.DB, embeds *[]discord.Embed) {
	entries, err := services.FetchShop()
	if err != nil {
		utils.Log("An error occurred while fetching shop entries: " + err.Error())
		return
	}

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

	embed := services.DefaultEmbed("shop_changed", fmt.Sprintf("%s/echoppe", frontBaseUrl))
	namesField := &discord.EmbedField{
		Name:   "Items",
		Value:  "",
		Inline: true,
	}
	raritiesField := &discord.EmbedField{
		Name:   "Raretés",
		Value:  "",
		Inline: true,
	}
	pricesField := &discord.EmbedField{
		Name:   "Prix",
		Value:  "",
		Inline: true,
	}

	for _, entry := range entries {
		golden := ""
		if entry.ShopInventory.Golden {
			golden = " (dorée)"
		}

		namesField.Value += fmt.Sprintf("`%s%s`\n", entry.ShopInventory.Item.Name, golden)
		raritiesField.Value += fmt.Sprintf("`%s`\n", strings.Repeat("★", entry.ShopInventory.Item.Rarity))
		pricesField.Value += fmt.Sprintf("`%d`\n", entry.ShopInventory.Balance)
	}

	embed.Fields = []*discord.EmbedField{namesField, raritiesField, pricesField}

	return embed
}
