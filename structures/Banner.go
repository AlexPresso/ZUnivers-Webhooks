package structures

import "gorm.io/gorm"

type BannerInventoryEntry struct {
	Banner *Banner `json:"banner"`
}

type Banner struct {
	gorm.Model

	BannerID string `json:"id"`
	Title    string `json:"title" zu:"display=Titre"`
	Name     string `json:"name" zu:"display=Nom"`
	ImageURL string `json:"imageUrl" zu:"imageUrl=%s"`
}
