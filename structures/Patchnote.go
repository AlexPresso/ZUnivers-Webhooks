package structures

import (
	"gorm.io/gorm"
)

type Patchnote struct {
	gorm.Model

	PatchnoteID string      `json:"id"`
	Title       string      `json:"title" display:"Titre"`
	Slug        string      `json:"slug"`
	CreatedBy   string      `json:"createdBy" display:"Auteur"`
	Date        *CustomTime `json:"date"`
	ImageUrl    string      `json:"imageUrl"`
}
