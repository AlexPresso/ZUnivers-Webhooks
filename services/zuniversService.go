package services

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"net/url"
)

func FetchConfigs() (configs []structures.Config, err error) {
	err = utils.Request("/public/recycle/config", "GET", nil, &configs)
	return
}

func FetchCurrentSeason() (season structures.Season, err error) {
	err = utils.Request("/public/tower/season", "GET", nil, &season)
	return
}

func FetchStatus() (status structures.Status, err error) {
	err = utils.Request("/app/status", "GET", nil, &status)
	return
}

func FetchItems() (items []structures.Item, err error) {
	err = utils.Request("/public/item", "GET", nil, &items)
	return
}

func FetchPacks() (packs []structures.Pack, err error) {
	err = utils.Request("/public/pack", "GET", nil, &packs)
	return
}

func FetchPatchnotes() (patchnotes []structures.Patchnote, err error) {
	err = utils.Request("/public/post", "GET", nil, &patchnotes)
	return
}

func FetchBanners() (banners []structures.BannerInventoryEntry, err error) {
	err = utils.Request("/public/banner/ZeratoR%231337", "GET", nil, &banners)
	return
}

func FetchEvents() (events []structures.Event, err error) {
	err = utils.Request("/public/event", "GET", nil, &events)
	return
}

func FetchUserDetail(discordTag string) (detail structures.UserDetail, err error) {
	err = utils.Request("/public/user/"+url.QueryEscape(discordTag), "GET", nil, &detail)
	return
}

func FetchAchievementCategories() (categories []structures.AchievementCategory, err error) {
	detail, err := FetchUserDetail("Alex'Presso#5480")
	return detail.AchievementCategories, err
}

func FetchAchievements(categoryId string) (achProgress []structures.AchievementProgress, err error) {
	err = utils.Request("/public/achievement/Alex'Presso%235480/"+categoryId, "GET", nil, &achProgress)
	return
}

func FetchChallenges() (chProgress []structures.ChallengeProgress, err error) {
	err = utils.Request("/public/challenge", "GET", nil, &chProgress)
	return
}
