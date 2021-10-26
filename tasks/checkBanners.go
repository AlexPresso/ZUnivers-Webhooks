package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkBanners(db *gorm.DB) {
	bannerEntries, err := services.FetchBanners()
	if err != nil {
		utils.Log("An error occurred while fetching bannerEntries: " + err.Error())
		return
	}

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
				services.DispatchEvent("banner_changed", *dbBanner, *banner)
			}
		} else if len(dbBanners) > 0 {
			services.DispatchEvent("new_banner", nil, *banner)
		}

		banners = append(banners, banner)
	}

	db.Save(&banners)
}
