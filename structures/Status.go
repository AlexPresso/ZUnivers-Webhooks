package structures

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model

	ApplicationVersion string `json:"applicationVersion" zu:"display=Version"`
	DbVersion          string `json:"dbVersion" zu:"display=Version de la BDD"`
	InstanceId         string `json:"instanceId"`
	Uptime             int    `json:"uptime"`
}
