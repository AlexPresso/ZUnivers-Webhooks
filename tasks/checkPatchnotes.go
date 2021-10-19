package tasks

import (
	"gorm.io/gorm"
	"time"
	"zunivers-webhooks/services"
	"zunivers-webhooks/structures"
	"zunivers-webhooks/utils"
)

func checkPatchnotes(db *gorm.DB) {
	patchnotes, err := services.FetchPatchnotes()
	if err != nil {
		utils.Log("An error occurred while fetching patch notes: " + err.Error())
		return
	}

	if len(patchnotes) == 0 {
		return
	}

	var latestPatchnote structures.Patchnote
	if res := db.Last(&latestPatchnote); res.Error == nil {
		if time.Time(*latestPatchnote.Date).Before(time.Time(*patchnotes[0].Date)) {
			services.DispatchEvent("new_patchnote", nil, patchnotes[0])
			db.Save(&patchnotes[0])
		}
	} else {
		db.Save(&patchnotes[0])
	}
}
