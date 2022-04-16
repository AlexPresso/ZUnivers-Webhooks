package structures

import "gorm.io/gorm"

type Item struct {
	gorm.Model

	ItemID string `json:"id" zu:"imageUrl=/item/%s-false"`
	Genre  string `json:"genre" zu:"display=Genre"`
	Name   string `json:"name" zu:"display=Nom"`
	Rarity uint32 `json:"rarity" zu:"display=Rareté"`
	Slug   string `json:"slug" zu:"url=/carte/%s"`
}

type Pack struct {
	gorm.Model

	PackID string `json:"id"`
	Name   string `json:"name" zu:"display=Nom;url=/catalogue/%s"`
	Year   uint32 `json:"year" zu:"display=Année"`
}
