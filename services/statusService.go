package services

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
)

func FetchStatus() (status structures.Status, err error) {
	err = utils.Request("/app/status", "GET", nil, &status)
	return
}
