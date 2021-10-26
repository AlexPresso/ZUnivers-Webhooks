package structures

import (
	"gorm.io/gorm"
)

type Patchnote struct {
	gorm.Model

	PatchnoteID string    `json:"id"`
	Title       string    `json:"title" zu:"display=Titre"`
	CreatedBy   string    `json:"createdBy" zu:"display=Auteur"`
	Date        *DateTime `json:"date" zu:"display=Date"`
	ImageUrl    string    `json:"imageUrl" zu:"imageUrl=%s"`
	Slug        string    `json:"slug" zu:"url=/post/%s"`
}
