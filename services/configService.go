package services

import (
	"zunivers-webhooks/structures"
	"zunivers-webhooks/utils"
)

func FetchConfigs() (configs []structures.Config, err error) {
	err = utils.Request("/public/recycle/config", "GET", nil, &configs)
	return
}
