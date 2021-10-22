package structures

import (
	"gorm.io/gorm"
)

type Patchnote struct {
	gorm.Model

	PatchnoteID string    `json:"id"`
	Title       string    `json:"title" display:"Titre"`
	Slug        string    `json:"slug"`
	CreatedBy   string    `json:"createdBy" display:"Auteur"`
	Date        *DateTime `json:"date" display:"Date"`
	ImageUrl    string    `json:"imageUrl"`
}
