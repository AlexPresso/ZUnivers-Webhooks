package tasks

import (
	"fmt"
	"gorm.io/gorm"
	"zunivers-webhooks/services"
	"zunivers-webhooks/structures"
	"zunivers-webhooks/utils"
)

func checkConfigs(db *gorm.DB) {
	configs, err := services.FetchConfigs()
	if err != nil {
		utils.Log("An error occurred while fetching configs: " + err.Error())
		return
	}

	var dbConfigs []structures.Config
	db.Find(&dbConfigs)
	dbConfigMap := make(map[string]*structures.Config)
	for _, config := range dbConfigs {
		config := config
		dbConfigMap[fmt.Sprintf("%d:%t", config.Rarity, config.IsGolden)] = &config
	}

	for i := 0; i < len(configs); i++ {
		config := &configs[i]
		dbConfig := dbConfigMap[fmt.Sprintf("%d:%t", config.Rarity, config.IsGolden)]

		if dbConfig != nil {
			config.ID = dbConfig.ID

			if dbConfig.CraftValue != config.CraftValue || dbConfig.RecycleValue != config.RecycleValue {
				services.DispatchEvent("config_changed", *dbConfig, *config)
			}
		} else if len(dbConfigs) > 0 {
			services.DispatchEvent("config_changed", nil, *config)
		}
	}

	db.Save(&configs)
}
