package structures

import "gorm.io/gorm"

type Item struct {
	gorm.Model

	ItemID string `json:"id"`
	Genre  string `json:"genre" display:"Genre"`
	Name   string `json:"name" display:"Nom"`
	Rarity uint32 `json:"rarity" display:"Rareté"`
}

type Pack struct {
	gorm.Model

	PackID string `json:"id"`
	Name   string `json:"name" display:"Nom"`
}
