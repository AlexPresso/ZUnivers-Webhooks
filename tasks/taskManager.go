package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"time"
)

func ScheduleTasks(db *gorm.DB) {
	s := gocron.NewScheduler(time.Local)

	_, _ = s.Every(20).Minutes().Do(checkInfos, db)
	_, _ = s.Every(1).Days().At("00:01").Do(newDay, db)

	s.StartBlocking()
}

func checkInfos(db *gorm.DB) {
	checkStatus(db)
	checkConfigs(db)
	checkPatchnotes(db)
	checkItems(db)
	checkBanners(db)

	utils.Log("Checked for infos.")
}

func newDay(db *gorm.DB) {
	services.DispatchEvent("new_day", nil, nil)
	checkSeason(db)
}
