package structures

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model

	ApplicationVersion string `json:"applicationVersion" display:"Version"`
	DbVersion          string `json:"dbVersion" display:"Version de la BDD"`
	InstanceId         string `json:"instanceId"`
	Uptime             int    `json:"uptime"`
}
