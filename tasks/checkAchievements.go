package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkAchievementCategories(db *gorm.DB, embeds *[]discord.Embed) {
	categories, err := services.FetchAchievementCategories()
	if err != nil {
		utils.Log("An error occured while fetching achievement categories: " + err.Error())
		return
	}

	var dbCategories []structures.AchievementCategory
	db.Find(&dbCategories)
	dbCategoriesMap := make(map[string]*structures.AchievementCategory)
	for _, category := range dbCategories {
		category := category
		dbCategoriesMap[category.CategoryID] = &category
	}

	var dbAchievements []structures.Achievement
	db.Find(&dbAchievements)
	dbAchievementsMap := make(map[string]*structures.Achievement)
	for _, achievement := range dbAchievements {
		achievement := achievement
		dbAchievementsMap[achievement.AchievementID] = &achievement
	}

	for i := 0; i < len(categories); i++ {
		category := &categories[i]
		dbCategory := dbCategoriesMap[category.CategoryID]

		if dbCategory != nil {
			category.ID = dbCategory.ID

			if utils.AreDifferent(*category, *dbCategory) {
				*embeds = append(*embeds, *services.MakeEmbed("achievement_category_changed", *dbCategory, *category))
			}
		} else if len(dbCategories) > 0 {
			*embeds = append(*embeds, *services.MakeEmbed("new_achievement_category", nil, *category))
		}

		checkAchievements(category.CategoryID, db, embeds, dbAchievementsMap)
	}

	db.Save(&categories)
}

func checkAchievements(categoryId string, db *gorm.DB, embeds *[]discord.Embed, dbAchievementsMap map[string]*structures.Achievement) {
	achProgress, err := services.FetchAchievements(categoryId)
	if err != nil {
		utils.Log("An error occured while fetching achievements: " + err.Error())
		return
	}

	var achievements []*structures.Achievement

	for i := 0; i < len(achProgress); i++ {
		achievement := &achProgress[i].Achievement
		achievements = append(achievements, *achievement)
		dbAchievement := dbAchievementsMap[(*achievement).AchievementID]

		if dbAchievement != nil {
			(*achievement).ID = dbAchievement.ID

			if utils.AreDifferent(**achievement, *dbAchievement) {
				*embeds = append(*embeds, *services.MakeEmbed("achievement_changed", *dbAchievement, **achievement))
			}
		} else if len(dbAchievementsMap) > 0 {
			*embeds = append(*embeds, *services.MakeEmbed("new_achievement", nil, **achievement))
		}
	}

	db.Save(&achievements)
}
