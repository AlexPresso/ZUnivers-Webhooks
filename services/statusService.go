package services

import (
	"zunivers-webhooks/structures"
	"zunivers-webhooks/utils"
)

func FetchStatus() (status structures.Status, err error) {
	err = utils.Request("/app/status", "GET", nil, &status)
	return
}
