package tasks

import (
	"github.com/alexpresso/zunivers-webhooks/services"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"gorm.io/gorm"
)

func checkStatus(db *gorm.DB) {
	status, err := services.FetchStatus()
	if err != nil {
		utils.Log("An error occurred while fetching status: " + err.Error())
		return
	}

	var currStatus structures.Status
	if res := db.First(&currStatus); res.Error == nil {
		if currStatus.ApplicationVersion != status.ApplicationVersion || currStatus.DbVersion != status.DbVersion {
			services.DispatchEvent("status_changed", currStatus, status)
		}

		status.ID = currStatus.ID
	}

	db.Save(&status)
}
