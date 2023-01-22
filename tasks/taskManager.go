package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"time"
)

func ScheduleTasks(db *gorm.DB) {
	s := gocron.NewScheduler(time.Local)

	_, _ = s.Every(20).Minutes().Do(checkInfos, db)
	_, _ = s.Every(1).Minutes().Do(checkShop, db)
	_, _ = s.Every(1).Days().At("00:01").Do(newDay)

	s.StartBlocking()
}

func checkInfos(db *gorm.DB) {
	embeds := &[]discord.Embed{}

	checkStatus(db, embeds)
	checkConfigs(db, embeds)
	checkPatchnotes(db, embeds)
	checkItems(db, embeds)
	checkBanners(db, embeds)
	checkEvents(db, embeds)
	checkAchievementCategories(db, embeds)
	checkSeason(db, embeds)
	checkChallenges(db, embeds)

	utils.Log("Checked for infos.")

	services.DispatchEmbeds(embeds)
}

func checkShop(db *gorm.DB) {
	embeds := &[]discord.Embed{}
	checkShopEntries(db, embeds)

	utils.Log("Checked for shop entries.")

	services.DispatchEmbeds(embeds)
}

func newDay() {
	utils.Log("New day")

	services.DispatchEmbeds(&[]discord.Embed{
		*services.MakeEmbed("new_day", nil, nil),
	})
}
