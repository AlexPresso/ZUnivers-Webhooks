package tasks

import (
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const ConfigChangedEvent = "config_changed"

func checkConfigs(db *gorm.DB, embeds *[]discord.Embed) {
	if utils.EventsAllDisabled([]string{ConfigChangedEvent}) {
		return
	}

	configs, resSpec, err := services.FetchConfigs()
	if err != nil {
		utils.Log("An error occurred while fetching configs: " + err.Error())
		return
	}

	checkResponse(db, embeds, resSpec)

	var dbConfigs []structures.Config
	db.Find(&dbConfigs)
	dbConfigMap := make(map[string]*structures.Config)
	for _, config := range dbConfigs {
		config := config
		dbConfigMap[fmt.Sprintf("%d:%d:%t", config.ShinyLevel, config.Rarity, config.IsGolden)] = &config
	}

	for i := 0; i < len(configs); i++ {
		config := &configs[i]
		dbConfig := dbConfigMap[fmt.Sprintf("%d:%d:%t", config.ShinyLevel, config.Rarity, config.IsGolden)]

		if dbConfig != nil {
			config.ID = dbConfig.ID

			if utils.AreDifferent(*dbConfig, *config) {
				services.MakeEmbed(ConfigChangedEvent, *dbConfig, *config, embeds)
			}
		} else if len(dbConfigs) > 0 {
			services.MakeEmbed(ConfigChangedEvent, nil, *config, embeds)
		}
	}

	db.Save(&configs)
}
