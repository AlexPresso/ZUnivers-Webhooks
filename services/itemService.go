package services

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
)

func FetchItems() (items []structures.Item, err error) {
	err = utils.Request("/public/item", "GET", nil, &items)
	return
}

func FetchPacks() (packs []structures.Pack, err error) {
	err = utils.Request("/public/pack", "GET", nil, &packs)
	return
}
