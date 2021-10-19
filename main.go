package main

import (
	"zunivers-webhooks/database"
	"zunivers-webhooks/tasks"
	"zunivers-webhooks/utils"
)

func main() {
	utils.LoadConfig()
	db := database.Init()
	tasks.ScheduleTasks(db)
}
