package services

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"image"
	"net/url"
)

func FetchConfigs() (configs []structures.Config, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/recycle/config", "GET", nil, &configs, &resSpec)
	return
}

func FetchCurrentSeason() (season structures.Season, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/tower/season", "GET", nil, &season, &resSpec)
	return
}

func FetchStatus() (status structures.Status, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/app/status", "GET", nil, &status, &resSpec)
	return
}

func FetchItems() (items []structures.Item, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/item", "GET", nil, &items, &resSpec)
	return
}

func FetchPacks() (packs []structures.Pack, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/pack", "GET", nil, &packs, &resSpec)
	return
}

func FetchPatchnotes() (patchnotes []structures.Patchnote, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/post", "GET", nil, &patchnotes, &resSpec)
	return
}

func FetchBanners() (banners []structures.BannerInventoryEntry, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/banner/zerator", "GET", nil, &banners, &resSpec)
	return
}

func FetchEvents() (events []structures.Event, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/event/current", "GET", nil, &events, &resSpec)
	return
}

func FetchUserDetail(discordTag string) (detail structures.UserDetail, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/user/"+url.QueryEscape(discordTag), "GET", nil, &detail, &resSpec)
	return
}

func FetchAchievementCategories() (categories []structures.AchievementCategory, resSpec structures.JsonResponseSpec, err error) {
	detail, resSpec, err := FetchUserDetail("alexpresso")
	return detail.AchievementCategories, resSpec, err
}

func FetchAchievements(categoryId string) (achProgress []structures.AchievementProgress, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/achievement/alexpresso/"+categoryId, "GET", nil, &achProgress, &resSpec)
	return
}

func FetchChallenges() (chProgress []structures.ChallengeProgress, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/challenge", "GET", nil, &chProgress, &resSpec)
	return
}

func FetchShop() (entries []structures.ShopEntry, resSpec structures.JsonResponseSpec, err error) {
	err = utils.Request("/public/shop", "GET", nil, &entries, &resSpec)
	return
}

func FetchLogo() (logo image.Image, err error) {
	err = utils.Request("/assets/logo-mini.png", "GET", nil, &logo, &structures.JsonResponseSpec{})
	return
}
