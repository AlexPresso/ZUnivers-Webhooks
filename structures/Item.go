package structures

import "gorm.io/gorm"

type Item struct {
	gorm.Model

	ItemID string   `json:"id" imageUrl:"/item/%s-false"`
	Genre  string   `json:"genre" display:"Genre"`
	Name   string   `json:"name" display:"Nom"`
	Rarity uint32   `json:"rarity" display:"Rareté"`
	URLs   []string `json:"urls"`
}

type Pack struct {
	gorm.Model

	PackID string `json:"id"`
	Name   string `json:"name" display:"Nom"`
}
