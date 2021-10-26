package structures

import "gorm.io/gorm"

type BannerInventoryEntry struct {
	Banner *Banner `json:"banner"`
}

type Banner struct {
	gorm.Model

	BannerID string `json:"id"`
	Title    string `json:"title" display:"Titre"`
	Name     string `json:"name" display:"Nom"`
	ImageURL string `json:"imageUrl" imageUrl:"%s;image"`
}
