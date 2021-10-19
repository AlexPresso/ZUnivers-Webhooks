package services

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/alexpresso/zunivers-webhooks/utils"
)

func FetchPatchnotes() (patchnotes []structures.Patchnote, err error) {
	err = utils.Request("/public/post", "GET", nil, &patchnotes)
	return
}
