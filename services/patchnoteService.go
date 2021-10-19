package services

import (
	"zunivers-webhooks/structures"
	"zunivers-webhooks/utils"
)

func FetchPatchnotes() (patchnotes []structures.Patchnote, err error) {
	err = utils.Request("/public/post", "GET", nil, &patchnotes)
	return
}
