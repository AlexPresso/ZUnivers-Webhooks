package services

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
)

func FetchConfigs() (configs []structures.Config, err error) {
	err = utils.Request("/public/recycle/config", "GET", nil, &configs)
	return
}
