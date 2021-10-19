package main

import (
	"github.com/alexpresso/zunivers-webhooks/database"
	"github.com/alexpresso/zunivers-webhooks/tasks"
	"github.com/alexpresso/zunivers-webhooks/utils"
)

func main() {
	utils.LoadConfig()
	db := database.Init()
	tasks.ScheduleTasks(db)
}
