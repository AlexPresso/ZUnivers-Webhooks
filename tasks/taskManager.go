package tasks

import (
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"time"
	"zunivers-webhooks/utils"
)

func ScheduleTasks(db *gorm.DB) {
	s := gocron.NewScheduler(time.Local)

	_, _ = s.Every(20).Minutes().Do(checkInfos, db)
	_, _ = s.Every(1).Days().At("00:01").Do(newDay)

	s.StartBlocking()
}

func checkInfos(db *gorm.DB) {
	checkStatus(db)
	checkConfigs(db)
	checkPatchnotes(db)

	utils.Log("Checked for infos.")
}
