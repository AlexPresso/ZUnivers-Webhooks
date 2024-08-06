package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

const NewBannerEvent = "new_banner"
const BannerChangedEvent = "banner_changed"

func checkBanners(db *gorm.DB, embeds *[]discord.Embed) {
	if !utils.EventsEnabled([]string{NewBannerEvent, BannerChangedEvent}) {
		return
	}

	bannerEntries, resSpec, err := services.FetchBanners()
	if err != nil {
		utils.Log("An error occurred while fetching bannerEntries: " + err.Error())
		return
	}

	checkResponse(db, embeds, BannerChangedEvent, resSpec)

	var dbBanners []structures.Banner
	db.Find(&dbBanners)
	bannersMap := make(map[string]*structures.Banner)
	for _, banner := range dbBanners {
		banner := banner
		bannersMap[banner.BannerID] = &banner
	}

	var banners []*structures.Banner
	for i := 0; i < len(bannerEntries); i++ {
		banner := bannerEntries[i].Banner
		dbBanner := bannersMap[banner.BannerID]

		if dbBanner != nil {
			banner.ID = dbBanner.ID

			if utils.AreDifferent(*dbBanner, *banner) {
				*embeds = append(*embeds, *services.MakeEmbed(BannerChangedEvent, *dbBanner, *banner))
			}
		} else if len(dbBanners) > 0 {
			*embeds = append(*embeds, *services.MakeEmbed(NewBannerEvent, nil, *banner))
		}

		banners = append(banners, banner)
	}

	db.Save(&banners)
}
